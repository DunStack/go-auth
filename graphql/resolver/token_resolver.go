package resolver

import (
	"github.com/dunstack/go-auth"
	"github.com/dunstack/go-auth/model/identity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/graph-gophers/graphql-go"
)

type tokenResolver struct {
	identity *identity.Identity
	idToken  *jwt.Token
	config   *auth.Config
}

func (tokenResolver) Type() string {
	return "Bearer"
}

func (r tokenResolver) ExpiresAt() (graphql.Time, error) {
	t, err := r.idToken.Claims.GetExpirationTime()
	if err != nil {
		return graphql.Time{}, err
	}
	return graphql.Time{Time: t.Time}, nil
}

func (r tokenResolver) IDToken() (string, error) {
	privKey, err := r.config.Token.IDToken.PrivateKey()
	if err != nil {
		return "", err
	}
	return r.idToken.SignedString(privKey)
}
