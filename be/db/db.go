package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/Armatorix/GoPress/be/db/ent/ent"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type Config struct {
	DSN               string `env:"DATABASE_URL,required" envDefault:"user=postgres host=localhost port=5432 dbname=postgres"`
	DataEncryptionKey string `env:"DB_DATA_ENCRYPTION_KEY,required" validate:"len=32"`
}

type DB struct {
	db *sql.DB
	ent.ClientTx
}

func NewWithMigrate(cfg Config) (*DB, error) {
	sqltrace.Register("pgx", stdlib.GetDefaultDriver())
	db, err := sqltrace.Open("pgx", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("open db failed: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping failed: %w", err)
	}
	db.SetMaxIdleConns(21)
	db.SetMaxOpenConns(37)

	goose.SetBaseFS(migrationsFS)

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, fmt.Errorf("database: can't set dialect %w", err)
	}

	if err := goose.Up(db, "migrations", goose.WithAllowMissing()); err != nil {
		return nil, fmt.Errorf("migrations up failed: %w", err)
	}

	client := ent.NewClient(
		ent.Driver(
			entsql.OpenDB(
				dialect.Postgres, db,
			),
		),
	)

	return &DB{
		db:       db,
		ClientTx: client,
	}, nil
}

func (db *DB) Close() error {
	if err := db.db.Close(); err != nil {
		return fmt.Errorf("database: can't close db: %w", err)
	}
	if c, ok := db.ClientTx.(*ent.Client); ok {
		if err := c.Close(); err != nil {
			return fmt.Errorf("database: can't close ent client: %w", err)
		}
	}
	return nil
}

func (db *DB) ExecuteInTx(ctx context.Context, exec func(context.Context, *DB) error) error {
	// if someone calls this method on a tx, we just execute the function
	// and assume that the caller is responsible for committing/rolling back
	// current transaction
	if _, ok := db.ClientTx.(*ent.Tx); ok {
		return exec(ctx, db)
	}

	// otherwise execute this function on db with a new transaction
	tx, err := db.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := exec(ctx, tx); err != nil {
		return err
	}

	return tx.Commit()
}

func (db *DB) BeginTx(ctx context.Context) (*DB, error) {
	tx, err := db.ClientTx.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("database: can't begin tx: %w", err)
	}
	return &DB{ClientTx: tx}, nil
}

func (db *DB) Rollback() {
	_ = db.ClientTx.Rollback()
}
