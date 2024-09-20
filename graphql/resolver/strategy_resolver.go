package resolver

import (
	"github.com/dunstack/go-auth/strategy"
)

type strategyResolver struct {
	strategy.Strategy
}

func (r strategyResolver) ToStrategyOAuth() (*strategy.StrategyOAuth, bool) {
	s, ok := r.Strategy.(*strategy.StrategyOAuth)
	return s, ok

}
