package database

import (
	"embed"
	"errors"
	"fmt"
	"time"
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/lembata/para/pkg/logger"
	//"github.com/mattn/go-sqlite3"
)

var logger = log.NewLogger()
var appSchemaVersion = uint(1)

//go:embed migrations/*.sql
var migrationsBox embed.FS

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

func NewDatabase() *Database {
	return &Database{
		lockChan: make(chan struct{}, 1),
	}
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
	db.lock()
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
		return err
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

// Locking db for writing
func (db *Database) lock() {
	db.lockChan <- struct{}{}
}

// Unlocking db after writing
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

