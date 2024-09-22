package auth

import (
	"crypto/ed25519"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenConfig struct {
	IDToken *IDTokenConfig
}

type IDTokenConfig struct {
	Key      string
	Lifetime time.Duration

	privateKey ed25519.PrivateKey
}

func (c IDTokenConfig) NewToken(sub string) *jwt.Token {
	now := time.Now()
	return jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.RegisteredClaims{
		Subject:   sub,
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(c.Lifetime)),
	})
}

func (c *IDTokenConfig) PrivateKey() (ed25519.PrivateKey, error) {
	if c.privateKey == nil {
		seed, err := base64.RawStdEncoding.DecodeString(c.Key)
		if err != nil {
			return nil, err
		}
		c.privateKey = ed25519.NewKeyFromSeed(seed)
	}
	return c.privateKey, nil
}
