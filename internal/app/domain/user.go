package domain

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID   string
	Name string
}

func NewUser(name string) (*User, error) {
	id := uuid.Must(uuid.NewV4(), nil).String()
	user := &User{
		ID:   id,
		Name: name,
	}
	return user, nil
}
