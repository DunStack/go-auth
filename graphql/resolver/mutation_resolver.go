package resolver

import (
	"context"
	"database/sql"

	"github.com/dunstack/go-auth"
	"github.com/dunstack/go-auth/graphql/resolver/input"
	"github.com/dunstack/go-auth/model/credential"
	credentialPassword "github.com/dunstack/go-auth/model/credential_password"
	"github.com/dunstack/go-auth/model/identity"
	"github.com/dunstack/go-auth/strategy"
	"github.com/go-playground/validator/v10"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type mutationResolver struct {
	app      *auth.App
	validate *validator.Validate
}

func (r *mutationResolver) SignInWithPassword(
	ctx context.Context,
	args struct {
		Input input.PasswordInput
	},
) (*tokenResolver, error) {
	if err := r.validate.StructCtx(ctx, args.Input); err != nil {
		return nil, err
	}

	i := new(identity.Identity)
	query := r.app.DB().NewSelect().Model(i)

	if v := args.Input.Username; v != nil {
		query = query.Where("username = ?", v)
	}
	if v := args.Input.Email; v != nil {
		query = query.Where("email = ?", v)
	}
	if v := args.Input.Phone; v != nil {
		query = query.Where("phone = ?", v)
	}

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	cp := new(credentialPassword.CredentialPassword)
	if err := r.app.DB().NewSelect().Model(cp).
		Join("JOIN credentials AS c on c.credential_type = 'PASSWORD' AND c.credential_id = cp.id").
		Where("c.identity_id = ?", i.ID).
		Scan(ctx); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(cp.Password), []byte(args.Input.Password)); err != nil {
		return nil, err
	}

	return &tokenResolver{
		app:      r.app,
		identity: i,
	}, nil
}

func (r *mutationResolver) SignUp(
	ctx context.Context,
	args struct {
		Input input.PasswordInput
	},
) (*identityResolver, error) {
	if err := r.validate.StructCtx(ctx, args.Input); err != nil {
		return nil, err
	}

	i := args.Input.ToIdentity()
	cp := args.Input.ToCredentialPassword()

	if err := r.app.DB().RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().Model(i).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewInsert().Model(cp).Exec(ctx); err != nil {
			return err
		}

		c := &credential.Credential{
			IdentityId:     i.ID,
			CredentialType: strategy.StrategyTypePassword,
			CredentialID:   cp.ID,
		}
		if _, err := tx.NewInsert().Model(c).Exec(ctx); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &identityResolver{
		object: i,
	}, nil
}
