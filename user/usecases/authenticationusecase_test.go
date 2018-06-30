package usecases

import (
	"testing"

	"github.com/easerp/erp-core/user"
	uuid "github.com/satori/go.uuid"
)

type mockUserRepo struct{}

func (repo *mockUserRepo) Create(*user.User) error {
	return nil
}

func (repo *mockUserRepo) Update(*user.User) error {
	return nil
}

func (repo *mockUserRepo) Find(uuid.UUID) (usr *user.User, err error) {
	return getMockUser("admin", "admin@google.com")
}

func getMockUser(login, email string) (usr *user.User, err error) {
	userID, err := uuid.NewV4()
	if err != nil {
		return
	}

	hash, err := user.HashPassword("rahasiadong")
	usr = &user.User{
		userID,
		"admin",
		hash,
		"admin@google.com",
		"Admin",
		"Ganteng",
		true,
		false,
	}
	return
}

func (repo *mockUserRepo) Delete(uuid.UUID) error {
	return nil
}

func (repo *mockUserRepo) FindByEmail(email string) (*user.User, error) {
	return getMockUser("admin", email)
}

func (repo *mockUserRepo) FindByLogin(login string) (*user.User, error) {
	return getMockUser(login, "admin@google.com")
}

func TestAuthenticate(t *testing.T) {
	repo := &mockUserRepo{}

	uc := &authenticationUseCase{repo}

	usr, err := uc.Authenticate("admin", "rahasiadong")
	if err != nil {
		t.Error(err)
	}

	if usr.Login != "admin" {
		t.Errorf("Expected login: %s, actual: %s", "admin", usr.Login)
	}

	usr, err = uc.Authenticate("admin", "gombal")
	if err == nil {
		t.Errorf("Expecting error")
	}

	if err.Error() != "Invalid Credential" {
		t.Errorf("Expected error: %s, actual: %s", "Invalid Credential", err.Error())
	}
}
