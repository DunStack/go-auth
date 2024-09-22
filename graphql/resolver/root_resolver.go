package resolver

import (
	"github.com/dunstack/go-auth"
	"github.com/go-playground/validator/v10"
)

func NewRootResolver(config *auth.Config) *RootResolver {
	return &RootResolver{
		config:   config,
		validate: validator.New(),
	}
}

type RootResolver struct {
	config   *auth.Config
	validate *validator.Validate
}

func (r *RootResolver) Query() *queryResolver {
	return &queryResolver{}
}

func (r *RootResolver) Mutation() *mutationResolver {
	return &mutationResolver{
		config:   r.config,
		validate: r.validate,
	}
}
