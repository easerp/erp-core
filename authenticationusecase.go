package core

type AuthenticationError struct {
	err string
}

func (ae *AuthenticationError) Error() string {
	return ae.err
}

type AuthenticationUseCase interface {
	Authenticate(login, password string) (usr User, err error)
}

type authenticationUseCase struct {
	repo UserRepo
}

func (uc *authenticationUseCase) Authenticate(login, password string) (
	usr *User, err error) {

	usr, err = uc.repo.FindByLogin(login)
	if err != nil {
		err = &AuthenticationError{"Invalid Credential"}
		return
	}

	valid := CheckPasswordHash(password, usr.Password)
	if !valid {
		err = &AuthenticationError{"Invalid Credential"}
		return
	}

	return
}
