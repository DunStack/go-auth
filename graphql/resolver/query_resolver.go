package resolver

import "github.com/dunstack/go-auth"

type queryResolver struct {
	cfg *auth.Config
}

func (c *queryResolver) Strategies() []strategyResolver {
	var strategies []strategyResolver
	for _, s := range c.cfg.Strategies {
		strategies = append(strategies, strategyResolver{s})
	}
	return strategies
}
