package credential

import "github.com/dunstack/go-auth/strategy"

type Credential struct {
	ID             int `bun:",pk,autoincrement"`
	CredentialType strategy.StrategyType
	CredentialID   int
}
