package resolver

type RootResolver struct{}

func (*RootResolver) Query() *queryResolver {
	return &queryResolver{}
}
