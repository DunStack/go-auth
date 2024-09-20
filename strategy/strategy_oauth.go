package strategy

import (
	"golang.org/x/exp/maps"
	"golang.org/x/oauth2"
)

var _ Strategy = new(StrategyOAuth)

type StrategyOAuth map[string]*oauth2.Config

func (StrategyOAuth) Type() StrategyType {
	return StrategyTypeOAuth
}

func (s StrategyOAuth) Providers() []string {
	return maps.Keys(s)
}
