package graphql

import (
	"context"
	_ "embed"
	"net/http"

	"github.com/dunstack/go-auth"
	"github.com/dunstack/go-auth/graphql/resolver"
	"github.com/dunstack/grapher"
	"github.com/graph-gophers/graphql-go"
)

//go:embed schema.gql
var schemaString string

func NewHandler(app *auth.App, opts HandlerOptions) *GraphQLHandler {
	schema := graphql.MustParseSchema(schemaString, resolver.NewRootResolver(app))
	var grapherOpts []grapher.HandlerOption

	if e := opts.Explorer; e != "" {
		grapherOpts = append(grapherOpts,
			grapher.WithExplorer(e),
			grapher.WithContext(func(r *http.Request) context.Context {
				return &resolver.Context{
					Context: r.Context(),
					App:     app,
					Request: r,
				}
			}),
		)
	}

	return &GraphQLHandler{
		Handler: grapher.NewHandler(schema, grapherOpts...),
	}
}

type GraphQLHandler struct {
	*grapher.Handler
}

type HandlerOptions struct {
	Explorer string
}
