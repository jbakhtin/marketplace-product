package postgres

import (
	"database/sql"
	"embed"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
)

type Config interface {
	GetDbDriver() string
	GetDbHost() string
	GetDbPort() string
	GetDbName() string
	GetDbUser() string
	GetDbPassword() string
}

//go:embed migrations/*.sql
var EmbedMigrations embed.FS

func NewSQLClient(cfg Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.GetDbDriver(), getDbDSN(cfg))
	if err != nil {
		return nil, errors.Wrap(err, "db open")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "db ping")
	}

	goose.SetBaseFS(EmbedMigrations)

	err = goose.SetDialect("postgres")
	if err != nil {
		return nil, errors.Wrap(err, "set dialect")
	}

	err = goose.Up(db, "migrations")
	if err != nil {
		return nil, errors.Wrap(err, "run migrations")
	}

	return db, nil
}

func getDbDSN(cfg Config) string {
	return "postgres://" + cfg.GetDbUser() + ":" + cfg.GetDbPassword() + "@" + cfg.GetDbHost() + ":" + cfg.GetDbPort() + "/" + cfg.GetDbName()
}
