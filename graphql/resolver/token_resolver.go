package resolver

import "github.com/dunstack/go-auth/model/identity"

type tokenResolver struct {
	i *identity.Identity
}

func (tokenResolver) Type() string {
	return "Bearer"
}

func (r tokenResolver) IDToken() string {
	// TODO: generate JWT
	return ""
}
