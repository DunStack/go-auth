package auth

import (
	_ "embed"

	"github.com/dunstack/go-auth/resolver"
	"github.com/dunstack/grapher"
	"github.com/graph-gophers/graphql-go"
)

//go:embed schema.gql
var schemaString string

type GraphQLConfig struct {
	Explorer string
}

func newGraphQLHandler(cfg *GraphQLConfig) *GraphQLHandler {
	schema := graphql.MustParseSchema(schemaString, &resolver.RootResolver{})
	var opts []grapher.HandlerOption

	if e := cfg.Explorer; e != "" {
		opts = append(opts, grapher.WithExplorer(e))
	}

	return &GraphQLHandler{
		Handler: grapher.NewHandler(schema, opts...),
	}
}

type GraphQLHandler struct {
	*grapher.Handler
}
