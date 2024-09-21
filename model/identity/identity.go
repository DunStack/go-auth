package identity

type Identity struct {
	ID       int    `bun:",pk,autoincrement"`
	Username string `bun:",nullzero"`
	Email    string `bun:",nullzero"`
	Phone    string `bun:",nullzero"`
}
