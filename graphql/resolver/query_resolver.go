package resolver

import "context"

type queryResolver struct {
}

func (c *queryResolver) CurrentIdentity(ctx context.Context) (*identityResolver, error) {
	i, err := ctx.(*Context).Identity()
	if err != nil {
		return nil, err
	}
	return &identityResolver{object: i}, nil
}
