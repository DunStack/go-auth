package resolver

import (
	"strconv"

	"github.com/dunstack/go-auth/model/identity"
	"github.com/graph-gophers/graphql-go"
)

type identityResolver struct {
	object *identity.Identity
}

func (r *identityResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(r.object.ID))
}
