package credential

import "github.com/dunstack/go-auth"

type Credential struct {
	ID             int `bun:",pk,autoincrement"`
	IdentityId     int
	CredentialType auth.StrategyType
	CredentialID   int
}
