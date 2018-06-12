package user

import uuid "github.com/satori/go.uuid"

type User struct {
	ID        uuid.UUID
	Login     string
	Password  string
	Email     string
	FirstName string
	LastName  string
}

type UserRepo interface {
	Create(*User) error
	Update(*User) error
	Find(uuid.UUID) (*User, error)
	Delete(uuid.UUID) error
}
