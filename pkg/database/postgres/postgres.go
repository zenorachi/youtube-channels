package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

const (
	Driver       = "pgx"
	ConfigPrefix = "DB"
)

// DBConfig - configuration for the database.
// DSN - full database url (including user, password, etc.).
// MaxIdleConns - maximum number of idle (default: 100).
// MaxOpenConns - maximum number of open connections (default: 10).
type DBConfig struct {
	DSN            string
	MaxIdleConns   int    `split_words:"true" default:"100"`
	MaxOpenConns   int    `split_words:"true" default:"10"`
	MigrationTable string `split_words:"true"`
	MigrationDir   string `split_words:"true" default:"./migrations"`
	AutoMigrate    bool   `split_words:"true" default:"false"`
}

// NewDB creates a new DB connection using the given config DBConfig.
// Also runs auto migration if AutoMigrate in config value is true. Default path for migration directory is "./migrations".
func NewDB(cfg *DBConfig) (*sqlx.DB, error) {
	if cfg == nil {
		return nil, errors.New("postgres config is required")
	}

	pgxConfig, err := pgx.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, err
	}

	pgxConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	db := sqlx.NewDb(
		stdlib.OpenDB(*pgxConfig),
		Driver,
	)

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if !cfg.AutoMigrate {
		return db, nil
	}

	if err = goose.SetDialect(Driver); err != nil {
		return nil, err
	}

	goose.SetTableName(cfg.MigrationTable)
	if err = goose.Up(db.DB, cfg.MigrationDir); err != nil {
		return nil, err
	}

	return db, nil
}
