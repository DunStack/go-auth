package resolver

import (
	"context"

	"github.com/dunstack/go-auth"
)

type queryResolver struct {
	app *auth.App
}

func (c *queryResolver) Strategies() []strategyResolver {
	var strategies []strategyResolver
	for _, s := range c.app.Strategies {
		strategies = append(strategies, strategyResolver{s})
	}
	return strategies
}

func (c *queryResolver) CurrentIdentity(ctx context.Context) (*identityResolver, error) {
	i, err := ctx.(*Context).Identity()
	if err != nil {
		return nil, err
	}
	return &identityResolver{object: i}, nil
}
