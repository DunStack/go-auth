package resolver

import (
	"github.com/dunstack/go-auth"
	"github.com/go-playground/validator/v10"
)

func NewRootResolver(app *auth.App) *RootResolver {
	return &RootResolver{
		app:      app,
		validate: validator.New(),
	}
}

type RootResolver struct {
	app      *auth.App
	validate *validator.Validate
}

func (r *RootResolver) Query() *queryResolver {
	return &queryResolver{
		app: r.app,
	}
}

func (r *RootResolver) Mutation() *mutationResolver {
	return &mutationResolver{
		app:      r.app,
		validate: r.validate,
	}
}
