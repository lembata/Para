package database

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	sqlite3mig "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type Migrator struct {
	db      *Database
	migrate *migrate.Migrate
}

func NewMigrator(db *Database) (*Migrator, error) {
	m := &Migrator{
		db: db,
	}

	migrate, err := m.getMigrate()
	m.migrate = migrate

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Migrator) Migrate() error {
	m.db.schemaVersion = 0
	return nil
}

func (m *Migrator) Version() (uint, error) {
	v, _, error := m.migrate.Version()
	return uint(v), error
}

func (m *Migrator) RunMigration(ctx context.Context, newVersion uint) error {
	databaseSchemaVersion, _, _ := m.migrate.Version()

	if newVersion != databaseSchemaVersion+1 {
		return fmt.Errorf("invalid migration version %d, expected %d", newVersion, databaseSchemaVersion+1)
	}

	// run pre migrations as needed
	// if err := m.runCustomMigrations(ctx, preMigrations[newVersion]); err != nil {
	// 	return fmt.Errorf("running pre migrations for schema version %d: %w", newVersion, err)
	// }

	var err error
	if err = m.migrate.Up(); err != nil {
		// migration failed
		m.migrate.Down()
		return err
	}

	// run post migrations as needed
	// if err := m.runCustomMigrations(ctx, postMigrations[newVersion]); err != nil {
	// 	return fmt.Errorf("running post migrations for schema version %d: %w", newVersion, err)
	// }

	// update the schema version
	m.db.schemaVersion, _, _ = m.migrate.Version()

	return nil
}

func (m *Migrator) Close() {
	if m.migrate == nil {
		m.migrate.Close()
		m.migrate = nil
	}
}

func (m *Migrator) getLatestMigration() uint {
	return uint(9)
}

func (m *Migrator) getMigrate() (*migrate.Migrate, error) {
	migrations, err := iofs.New(migrationsBox, "migrations")
	if err != nil {
		return nil, err
	}

	const disableForeignKeys = true
	conn, err := m.db.open(disableForeignKeys)

	if err != nil {
		return nil, err
	}

	driver, err := sqlite3mig.WithInstance(conn.DB, &sqlite3mig.Config{})
	if err != nil {
		return nil, err
	}

	// use sqlite3Driver so that migration has access to durationToTinyInt
	return migrate.NewWithInstance(
		"iofs",
		migrations,
		m.db.dbPath,
		driver,
	)
}
