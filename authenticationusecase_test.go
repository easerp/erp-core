package core

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

type mockUserRepo struct{}

func (repo *mockUserRepo) Create(*User) error {
	return nil
}

func (repo *mockUserRepo) Update(*User) error {
	return nil
}

func (repo *mockUserRepo) Find(uuid.UUID) (usr *User, err error) {
	return getMockUser("admin", "admin@google.com")
}

func getMockUser(login, email string) (usr *User, err error) {
	userID, err := uuid.NewV4()
	if err != nil {
		return
	}

	hash, err := HashPassword("rahasiadong")
	usr = &User{
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

func (repo *mockUserRepo) FindByEmail(email string) (*User, error) {
	return getMockUser("admin", email)
}

func (repo *mockUserRepo) FindByLogin(login string) (*User, error) {
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
