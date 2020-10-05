package entities

import (
	"strings"
)

type User struct {
	id   int
	name string
}

func NewUser(
	id int,
	name string,
) *User {
	return &User{
		id:   id,
		name: name,
	}
}

func (u *User) ID() int      { return u.id }
func (u *User) Name() string { return u.name }

func (u *User) IsQualifyGreenMiles() bool {
	return strings.Contains(strings.ToLower(u.name), "green")
}
