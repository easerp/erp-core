package usecases

import (
	"github.com/easerp/erp-core/user"
)

type AuthenticationError struct {
	err string
}

func (ae *AuthenticationError) Error() string {
	return ae.err
}

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
		err = &AuthenticationError{"Invalid Credential"}
		return
	}

	valid := user.CheckPasswordHash(password, usr.Password)
	if !valid {
		err = &AuthenticationError{"Invalid Credential"}
		return
	}

	return
}
