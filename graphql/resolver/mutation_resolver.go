package resolver

import (
	"context"
	"database/sql"

	"github.com/dunstack/go-auth"
	"github.com/dunstack/go-auth/graphql/resolver/input"
	"github.com/dunstack/go-auth/model/credential"
	"github.com/dunstack/go-auth/strategy"
	"github.com/go-playground/validator/v10"
	"github.com/uptrace/bun"
)

type mutationResolver struct {
	app      *auth.App
	validate *validator.Validate
}

func (r *mutationResolver) SignUp(
	ctx context.Context,
	args struct {
		Input input.SignUpInput
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
