package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dunstack/go-auth"
	"github.com/dunstack/go-auth/graphql"
)

func main() {
	log.Println("start is starting ...")
	defer log.Println("server stopped")

	cfg := &auth.Config{
		DB: &auth.DBConfig{
			DSN: "postgresql://postgres:knlP5T1pK67c6JBd@laudably-lenient-fennec.data-1.use1.tembo.io:5432/postgres?sslmode=require",
		},
		Token: &auth.TokenConfig{
			IDToken: &auth.IDTokenConfig{
				Key:      "AyrDUek+N9jdhjpbadKBQNAX19GI5Qrh7/xCfMU62kE",
				Lifetime: 1 * time.Hour,
			},
		},
	}

	http.Handle("/graphql", graphql.NewHandler(cfg, graphql.HandlerOptions{
		Explorer: "static/graphiql.html",
	}))

	log.Println("server started at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
