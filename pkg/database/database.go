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
		return errors.New("database not initialized")
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

	if err != nil {
		return nil, fmt.Errorf("db.Open(): %w", err)
	}

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

func (db *Database) GetAccounts(ctx context.Context) (*sql.Rows, error) {
	logger.Debugf("Getting Accounts")

	/*
select a.id, a.name, sum(t.total) as transaction_amount
from accounts a
left join transactions t on t.account_id = a.id
group by a.id, a.name
	*/
	sqler := squirrel.Select("a.id", "a.name",
		"(a.opening_balance + ifnull(sum(t.total_amount), 0) - ifnull(sum(f.total_amount), 0) ) as delta").
		From("accounts a").
		LeftJoin("transactions f on f.from_account_id = a.id").
		LeftJoin("transactions t on t.to_account_id = a.id").
		// Columns("id", "name", "currency", "iban", "bic",
		// 	"account_number", "opening_balance",
		// 	"opening_balance_date", "notes",
		// 	"created_at", "updated_at",
		// 	"include_in_net_worth").
		GroupBy("a.id", "a.name", "a.opening_balance").
		Where("deleted = ?", false).
		OrderBy("a.id").
		Offset(0).
		Limit(100)
	
	return query(ctx, sqler)
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

	sql, args, err := stmt.ToSql()
	if err != nil {
		return nil, fmt.Errorf("generating sql: %w", err)
	}

	logger.Tracef("SQL: %s [%v]", sql, args)
	ret, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("executing `%s` [%v]: %w", sql, args, err)
	}

	return ret, nil
}

func query(ctx context.Context, stmt sqler) (*sql.Rows, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return nil, err
	}

	sql, args, err := stmt.ToSql()
	if err != nil {
		return nil, fmt.Errorf("generating sql: %w", err)
	}

	logger.Tracef("SQL: %s [%v]", sql, args)
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("executing `%s` [%v]: %w", sql, args, err)
	}

	return rows, nil
}



///////// transaction
