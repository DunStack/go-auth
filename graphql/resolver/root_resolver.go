package resolver

import "github.com/dunstack/go-auth"

type RootResolver struct {
	Config *auth.Config
}

func (r *RootResolver) Query() *queryResolver {
	return &queryResolver{
		cfg: r.Config,
	}
}
