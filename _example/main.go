package main

import (
	"crypto/ed25519"
	"log"
	"net/http"

	"github.com/dunstack/go-auth"
	"github.com/dunstack/go-auth/graphql"
	"github.com/dunstack/go-auth/strategy"
)

var seed = []byte{3, 42, 195, 81, 233, 62, 55, 216, 221, 134, 58, 91, 105, 210, 129, 64, 208, 23, 215, 209, 136, 229, 10, 225, 239, 252, 66, 124, 197, 58, 218, 65}

func main() {
	log.Println("start is starting ...")
	defer log.Println("server stopped")

	app := &auth.App{
		DSN:        "postgresql://postgres:knlP5T1pK67c6JBd@laudably-lenient-fennec.data-1.use1.tembo.io:5432/postgres?sslmode=require",
		PrivateKey: ed25519.NewKeyFromSeed(seed),
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
