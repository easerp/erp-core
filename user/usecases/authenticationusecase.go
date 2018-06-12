package usecases

import (
	"errors"

	"github.com/flexrp/flexrp.core/user"
)

type AuthenticationUseCase interface {
	Authenticate(login, password string) (usr *user.User, err error)
}

type authenticationUseCase struct {
	repo user.UserRepo
}

func (uc *authenticationUseCase) Authenticate(login, password string) (
	usr *user.User, err error) {

	usr, err = uc.repo.FindByLogin(login)
	if err != nil {
		err = errors.New("Invalid Credential")
		return
	}

	valid := user.CheckPasswordHash(password, usr.Password)
	if !valid {
		err = errors.New("Invalid Credential")
		return
	}

	return
}
