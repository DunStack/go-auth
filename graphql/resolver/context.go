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

func NewContext(config *auth.Config, request *http.Request) *Context {
	return &Context{
		Context: request.Context(),
		config:  config,
		request: request,
	}
}

type Context struct {
	context.Context
	config  *auth.Config
	request *http.Request

	identity *identity.Identity
}

func (c *Context) Identity() (*identity.Identity, error) {
	if c.identity == nil {
		scheme, tokenString, ok := strings.Cut(c.request.Header.Get("Authorization"), " ")
		if !ok {
			return nil, errors.New("invalid token")
		}
		if scheme != "Bearer" {
			return nil, errors.New("invalid scheme")
		}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			privKey, err := c.config.Token.IDToken.PrivateKey()
			if err != nil {
				return nil, err
			}
			return privKey.Public(), nil
		})
		if err != nil {
			return nil, err
		}

		id, err := token.Claims.GetSubject()
		if err != nil {
			return nil, err
		}

		i := new(identity.Identity)
		if err := c.config.DB.Client().NewSelect().Model(i).Where("id = ?", id).Scan(c); err != nil {
			return nil, err
		}
		c.identity = i
	}
	return c.identity, nil
}
