package core

import uuid "github.com/satori/go.uuid"

type User struct {
	ID          uuid.UUID
	Login       string
	Password    string
	Email       string
	FirstName   string
	LastName    string
	IsActivated bool
	IsSuspended bool
}

type UserRepo interface {
	Create(*User) error
	Update(*User) error
	Find(uuid.UUID) (*User, error)
	Delete(uuid.UUID) error

	FindByEmail(email string) (*User, error)
	FindByLogin(login string) (*User, error)
}
