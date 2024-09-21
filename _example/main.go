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

	cfg := &auth.Config{
		Strategies: []strategy.Strategy{
			&strategy.StrategyPassword{},
			&strategy.StrategyOAuth{
				"github": {},
			},
		},
	}
	http.Handle("/graphql", graphql.NewHandler(cfg, graphql.HandlerOptions{
		Explorer: "static/graphiql.html",
	}))

	log.Println("server started at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
