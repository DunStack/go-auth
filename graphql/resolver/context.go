package resolver

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/dunstack/go-auth"
	"github.com/dunstack/go-auth/model/identity"
	"github.com/golang-jwt/jwt/v5"
)

type Context struct {
	context.Context
	App     *auth.App
	Request *http.Request

	identity *identity.Identity
}

func (c *Context) Identity() (*identity.Identity, error) {
	if c.identity == nil {
		scheme, tokenString, ok := strings.Cut(c.Request.Header.Get("Authorization"), " ")
		if !ok {
			return nil, errors.New("invalid token")
		}
		if scheme != "Bearer" {
			return nil, errors.New("invalid scheme")
		}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return c.App.PrivateKey.Public(), nil
		})
		if err != nil {
			return nil, err
		}

		id, err := token.Claims.GetSubject()
		if err != nil {
			return nil, err
		}

		i := new(identity.Identity)
		if err := c.App.DB().NewSelect().Model(i).Where("id = ?", id).Scan(c); err != nil {
			return nil, err
		}
		c.identity = i
	}
	return c.identity, nil
}
