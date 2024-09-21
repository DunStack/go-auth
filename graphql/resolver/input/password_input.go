package input

import (
	credentialPassword "github.com/dunstack/go-auth/model/credential_password"
	"github.com/dunstack/go-auth/model/identity"
)

type PasswordInput struct {
	Username *string `validate:"required_without_all=Email Phone,omitempty"`
	Email    *string `validate:"required_without_all=Username Phone,omitempty,email"`
	Phone    *string `validate:"required_without_all=Username Email,omitempty,e164"`
	Password string
}

func (input PasswordInput) ToIdentity() *identity.Identity {
	i := new(identity.Identity)
	if v := input.Username; v != nil {
		i.Username = *v
	}
	if v := input.Email; v != nil {
		i.Email = *v
	}
	if v := input.Phone; v != nil {
		i.Phone = *v
	}
	return i
}
func (input PasswordInput) ToCredentialPassword() *credentialPassword.CredentialPassword {
	return &credentialPassword.CredentialPassword{
		Password: input.Password,
	}
}
