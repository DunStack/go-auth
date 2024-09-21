package credentialPassword

import (
	"context"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

var _ bun.BeforeAppendModelHook = new(CredentialPassword)

type CredentialPassword struct {
	bun.BaseModel `bun:"table:credentials_password,alias:cp"`

	ID       int `bun:",pk,autoincrement"`
	Password string
}

func (c *CredentialPassword) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery, *bun.UpdateQuery:
		return c.bscryptPassword()
	}
	return nil
}

func (c *CredentialPassword) bscryptPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	c.Password = string(bytes)
	return nil
}
