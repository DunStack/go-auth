package main

import (
	"log"
	"net/http"

	"github.com/dunstack/go-auth"
	"github.com/dunstack/go-auth/graphql"
	"github.com/dunstack/go-auth/strategy"
)

func main() {
	log.Println("start is starting ...")
	defer log.Println("server stopped")

	app := &auth.App{
		DSN: "postgresql://postgres:knlP5T1pK67c6JBd@laudably-lenient-fennec.data-1.use1.tembo.io:5432/postgres?sslmode=require",
		Strategies: []strategy.Strategy{
			&strategy.StrategyPassword{},
			&strategy.StrategyOAuth{
				"github": {},
			},
		},
	}
	http.Handle("/graphql", graphql.NewHandler(app, graphql.HandlerOptions{
		Explorer: "static/graphiql.html",
	}))

	log.Println("server started at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
