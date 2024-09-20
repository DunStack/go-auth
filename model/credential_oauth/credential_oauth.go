package credentialOAuth

import "github.com/uptrace/bun"

type CredentialOAuth struct {
	bun.BaseModel `bun:"table:credentials_oauth"`

	ID       int `bun:",pk,autoincrement"`
	Provider string
	UID      string
}
