package usecases

import (
	"github.com/easerp/erp-core/user"
	uuid "github.com/satori/go.uuid"
)

type RegisterUseCase interface {
	RegisterUser(login, password, email, firstname, lastname string) (*user.User, error)
	ActivateUser(email string) error
}

type registerUseCase struct {
	repo user.UserRepo
}

func NewRegisterUseCase(repo user.UserRepo) *registerUseCase {
	return &registerUseCase{
		repo,
	}
}

func (uc *registerUseCase) RegisterUser(login, password, email, firstname, lastname string) (
	usr *user.User, err error) {

	ID, err := uuid.NewV4()
	if err != nil {
		return
	}

	hashedPass, err := user.HashPassword(password)
	if err != nil {
		return
	}

	usr = &user.User{
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
