package auth

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DBConfig struct {
	DSN string

	db *bun.DB
}

func (c *DBConfig) Client() *bun.DB {
	if c.db == nil {
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(c.DSN)))
		c.db = bun.NewDB(sqldb, pgdialect.New())
	}
	return c.db
}
