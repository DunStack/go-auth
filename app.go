package auth

import (
	"crypto/ed25519"
	"database/sql"

	"github.com/dunstack/go-auth/strategy"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type App struct {
	DSN        string
	Strategies []strategy.Strategy
	PrivateKey ed25519.PrivateKey

	db *bun.DB
}

func (app *App) DB() *bun.DB {
	if app.db == nil {
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(app.DSN)))
		app.db = bun.NewDB(sqldb, pgdialect.New())
	}
	return app.db
}
