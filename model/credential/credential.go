package credential

type CredentialType string

const (
	CredentialTypeOAuth CredentialType = "OAUTH"
)

type Credential struct {
	ID             int `bun:",pk,autoincrement"`
	CredentialType CredentialType
	CredentialID   int
}
