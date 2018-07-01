package core

import (
	uuid "github.com/satori/go.uuid"
)

type RegisterUseCase interface {
	RegisterUser(login, password, email, firstname, lastname string) (User, error)
	ActivateUser(email string) error
}

type registerUseCase struct {
	repo UserRepo
}

func NewRegisterUseCase(repo UserRepo) *registerUseCase {
	return &registerUseCase{
		repo,
	}
}

func (uc *registerUseCase) RegisterUser(login, password, email, firstname, lastname string) (
	usr *User, err error) {

	ID, err := uuid.NewV4()
	if err != nil {
		return
	}

	hashedPass, err := HashPassword(password)
	if err != nil {
		return
	}

	usr = &User{
		ID,
		login,
		hashedPass,
		email,
		firstname,
		lastname,
		false,
		false,
	}

	err = uc.repo.Create(usr)
	return
}

func (uc *registerUseCase) ActivateUser(email string) (err error) {
	usr, err := uc.repo.FindByEmail(email)
	if err != nil {
		return
	}

	usr.IsActivated = true

	err = uc.repo.Update(usr)
	return
}
