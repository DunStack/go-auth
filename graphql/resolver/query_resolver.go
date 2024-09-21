package resolver

import "github.com/dunstack/go-auth"

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
