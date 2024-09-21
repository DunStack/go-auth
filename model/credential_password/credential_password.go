package credentialPassword

import (
	"context"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

var _ bun.BeforeInsertHook = new(CredentialPassword)
var _ bun.BeforeUpdateHook = new(CredentialPassword)

type CredentialPassword struct {
	bun.BaseModel `bun:"table:credentials_password"`

	ID       int `bun:",pk,autoincrement"`
	Password string
}

func (c *CredentialPassword) BeforeInsert(ctx context.Context, query *bun.InsertQuery) error {
	return c.bscryptPassword()
}

func (c *CredentialPassword) BeforeUpdate(ctx context.Context, query *bun.UpdateQuery) error {
	if c.Password == "" {
		return nil
	}
	return c.bscryptPassword()
}

func (c *CredentialPassword) bscryptPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	c.Password = string(bytes)
	return nil
}
