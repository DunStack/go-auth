package auth

import (
	"github.com/dunstack/go-auth/strategy"
)

type Config struct {
	Strategies []strategy.Strategy
}
