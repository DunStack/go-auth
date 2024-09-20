package main

import (
	"net/http"

	"github.com/dunstack/go-auth"
)

func main() {
	app := auth.NewApp(&auth.Config{})
	http.Handle("/graphql", app.NewGraphQLHandler(&auth.GraphQLConfig{
		Explorer: "static/graphiql.html",
	}))
	http.ListenAndServe(":8080", nil)
}
