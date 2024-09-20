package resolver

type queryResolver struct{}

func (*queryResolver) Hello() string {
	return "world"
}
