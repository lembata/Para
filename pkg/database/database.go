package database

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/Masterminds/squirrel"
	"github.com/lembata/para/internal/entities"
	"github.com/lembata/para/pkg/logger"
	//"github.com/mattn/go-sqlite3"
)

var logger = log.NewLogger()
var appSchemaVersion = uint(1)

//go:embed migrations/*.sql
var migrationsBox embed.FS

var instance *Database

const (
	// Number of database connections to use
	// The same value is used for both the maximum and idle limit,
	// to prevent opening connections on the fly which has a notieable performance penalty.
	// Fewer connections use less memory, more connections increase performance,
	// but have diminishing returns.
	// 10 was found to be a good tradeoff.
	dbConns = 10
	// Idle connection timeout, in seconds
	// Closes a connection after a period of inactivity, which saves on memory and
	// causes the sqlite -wal and -shm files to be automatically deleted.
	dbConnTimeout = 30
	sqlite3Drive  = "sqlite3"
)

var (
	ErrorNotFound       = errors.New("not found")
	ErrorNotInitialized = errors.New("not initialized")
)

type Database struct {
	dbPath        string
	db            *sqlx.DB
	schemaVersion uint
	lockChan      chan struct{}
}

func Init() *Database {
	if instance == nil {
		instance = newDatabase()
	}

	return instance
}

func GetInstance() *Database {
	return instance
}

func (db *Database) Ready() error {
	if db.db == nil {
		return ErrorNotInitialized
	}

	return nil
}

func (db *Database) Close() error {
	logger.Info("Closing database")
	return db.db.Close()
}

func (db *Database) Open(connectionString string) error {
	logger.Info("Opening database")
	db.lockNoCtx()
	defer db.unlock()

	db.dbPath = connectionString
	var err error

	db.schemaVersion, _ = db.getSchemaVersion()
	logger.Infof("Database schema version: %d", db.schemaVersion)

	if db.schemaVersion == 0 {
		logger.Info("Initial database migration")
		err = db.RunAllMigrations()
	}

	if err != nil {
		logger.Errorf("error while running migrations: %v", err)
		return err
	}

	if db.db == nil {
		if db.db, err = db.open(false); err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) RunAllMigrations() error {
	ctx := context.Background()

	m, err := NewMigrator(db)
	if err != nil {
		return err
	}
	defer m.Close()

	databaseSchemaVersion, _, _ := m.migrate.Version()
	stepNumber := appSchemaVersion - databaseSchemaVersion
	if stepNumber != 0 {
		logger.Infof("Migrating database from version %d to %d", databaseSchemaVersion, appSchemaVersion)

		// run each migration individually, and run custom migrations as needed
		var i uint = 1
		for ; i <= stepNumber; i++ {
			newVersion := databaseSchemaVersion + i
			if err := m.RunMigration(ctx, newVersion); err != nil {
				return err
			}
		}
	}

	// re-initialise the database
	const disableForeignKeys = false
	db.db, err = db.open(disableForeignKeys)
	if err != nil {
		return fmt.Errorf("re-initializing the database: %w", err)
	}

	// optimize database after migration
	//err = db.Optimise(ctx)
	// if err != nil {
	// 	logger.Warnf("error while performing post-migration optimisation: %v", err)
	// }

	return nil
}

func (db *Database) Exec(query string, args ...any) error {
	db.lockNoCtx()
	defer db.unlock()

	result, err := db.db.Exec(query, args...)

	if err != nil {
		logger.Errorf("error while executing query: %v", err)
		return err
	}

	logger.Infof("result: %v", result)
	return nil
}

// lock locks the database for writing.
// This method will block until the lock is acquired of the context is cancelled.
func (db *Database) lock(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case db.lockChan <- struct{}{}:
		return nil
	}
}

// lock locks the database for writing. This method will block until the lock is acquired.
func (db *Database) lockNoCtx() {
	db.lockChan <- struct{}{}
}

// unlock unlocks the database
func (db *Database) unlock() {
	select {
	case <-db.lockChan:
		return
	default:
		panic("database is not locked")
	}
}

func (db *Database) getSchemaVersion() (uint, error) {
	migrator, err := NewMigrator(db)

	if err != nil {
		return 0, err
	}
	defer migrator.Close()

	ret, err := migrator.Version()

	return ret, err
}

func (db *Database) open(disableForeignKeys bool) (*sqlx.DB, error) {
	// https://github.com/mattn/go-sqlite3
	url := "file:" + db.dbPath + "?_journal=WAL&_sync=NORMAL&_busy_timeout=50"

	if !disableForeignKeys {
		url += "&_fk=true"
	}

	conn, err := sqlx.Open(sqlite3Drive, url)

	if err != nil {
		logger.Errorf("sqlx.Open(): %v", err)
		return nil, fmt.Errorf("sqlx.Open(): %w", err)
	}

	conn.SetMaxOpenConns(dbConns)
	conn.SetMaxIdleConns(dbConns)
	conn.SetConnMaxIdleTime(dbConnTimeout * time.Second)

	return conn, nil
}

func (db *Database) CreateAccount(ctx context.Context, account entities.AccountEntity) (int64, error) {
	logger.Debugf("Creating account: %v", account)

	sqler := squirrel.Insert("accounts").
		Columns("name", "currency", "iban", "bic",
			"account_number", "opening_balance",
			"opening_balance_date", "notes",
			"created_at", "updated_at",
			"include_in_net_worth").
		Values(account.Name, account.Currency, account.IBAN, account.BIC,
			account.AccountNumber, account.OpeningBalance,
			account.OpeningBalanceDate, account.Notes,
			account.CreateAt, account.UpdateAt,
			account.IncludeInNetWorth)

	result, err := exec(ctx, sqler)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (db *Database) EditAccount(ctx context.Context, account entities.AccountEntity) (int64, error) {
	logger.Debugf("Creating account: %v", account)

	sqler := squirrel.Update("accounts").
		Set("name", account.Name).
		Set("currency", account.Currency).
		Set("iban", account.IBAN).
		Set("bic", account.BIC).
		Set("account_number", account.AccountNumber).
		Set("opening_balance", account.OpeningBalance).
		Set("opening_balance_date", account.OpeningBalanceDate).
		Set("notes", account.Notes).
		Set("created_at", account.CreateAt).
		Set("updated_at", account.UpdateAt).
		Set("include_in_net_worth", account.IncludeInNetWorth).
		Where("id = ?", account.Id)

	result, err := exec(ctx, sqler)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (db *Database) GetAccountById(ctx context.Context, id int64) (entities.AccountEntity, error) {
	logger.Debugf("Getting Accounts")

	sqler := squirrel.Select("id", "name", "currency", "iban", "bic",
		"account_number", "opening_balance",
		"opening_balance_date", "notes",
		"created_at", "updated_at",
		"include_in_net_worth").
		From("accounts").
		Where("id = ?", id).
		Limit(1)

	rows, err := query(ctx, sqler)

	if err != nil {
		return entities.AccountEntity{}, err
	}

	defer func() { _ = rows.Close() }()

	if rows.Next() {
		var row entities.AccountEntity

		if err := rows.Scan(&row.Id, &row.Name,
			&row.Currency, &row.IBAN, &row.BIC,
			&row.AccountNumber, &row.OpeningBalance,
			&row.OpeningBalanceDate, &row.Notes,
			&row.CreateAt, &row.UpdateAt,
			&row.IncludeInNetWorth); err != nil {
			logger.Errorf("Error %v", err)
			return row, err
		}

		return row, nil
	}

	return entities.AccountEntity{}, ErrorNotFound
}

func (db *Database) GetAccounts(ctx context.Context, offset uint64, limit uint64, order string) ([]entities.AccountRow, error) {
	logger.Debugf("Getting Accounts")

	sqler := squirrel.Select("a.id as id", "a.name as name",
		"a.currency as currency",
		"(a.opening_balance + ifnull(sum(t.total_amount), 0) - ifnull(sum(f.total_amount), 0) ) as balance").
		From("accounts a").
		LeftJoin("transactions f on f.from_account_id = a.id").
		LeftJoin("transactions t on t.to_account_id = a.id").
		GroupBy("a.id", "a.name", "a.currency", "a.opening_balance").
		Where("deleted = ?", false).
		OrderBy(order).
		Offset(offset).
		Limit(limit)

	rows, err := query(ctx, sqler)

	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()
	var accounts []entities.AccountRow

	for rows.Next() {
		//var alb Album
		var row entities.AccountRow

		if err := rows.Scan(&row.Id, &row.Name,
			&row.Balance.Currency, &row.Balance.Value); err != nil {
			logger.Errorf("Error %v", err)
			break
		}

		accounts = append(accounts, row)
	}

	return accounts, nil
}

func newDatabase() *Database {
	return &Database{
		lockChan: make(chan struct{}, 1),
	}
}

// func exex() ( sql.Result, error) {
// 	return nil, nil
// }

type sqler interface {
	ToSql() (string, []interface{}, error)
}

func exec(ctx context.Context, stmt sqler) (sql.Result, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return nil, err
	}

	query, args, err := stmt.ToSql()
	if err != nil {
		return nil, fmt.Errorf("generating sql: %w", err)
	}

	logger.Tracef("SQL: %s [%v]", query, args)
	ret, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing `%s` [%v]: %w", query, args, err)
	}

	return ret, nil
}

func query(ctx context.Context, stmt sqler) (*sql.Rows, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return nil, err
	}

	query, args, err := stmt.ToSql()
	if err != nil {
		return nil, fmt.Errorf("generating sql: %w", err)
	}

	logger.Tracef("SQL: %s [%v]", query, args)
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing `%s` [%v]: %w", query, args, err)
	}

	return rows, nil
}

///////// transaction
