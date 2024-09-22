package resolver

import (
	"strconv"
	"time"

	"github.com/dunstack/go-auth"
	"github.com/dunstack/go-auth/model/identity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/graph-gophers/graphql-go"
)

type tokenResolver struct {
	identity *identity.Identity
	app      *auth.App
}

func (tokenResolver) Type() string {
	return "Bearer"
}

func (r tokenResolver) ExpiresAt() graphql.Time {
	t := time.Now().Add(1 * time.Hour)
	return graphql.Time{Time: t}
}

func (r tokenResolver) IDToken() (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(r.identity.ID),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(1 * time.Hour)),
	})
	return token.SignedString(r.app.PrivateKey)
}
