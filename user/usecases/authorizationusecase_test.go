package usecases

import (
	"testing"

	"github.com/flexrp/flexrp.core/user"
	uuid "github.com/satori/go.uuid"
)

type mockRoleRepo struct{}

func (repo *mockRoleRepo) Create(*user.Role) error {
	return nil
}

func (repo *mockRoleRepo) Update(*user.Role) error {
	return nil
}

func (repo *mockRoleRepo) Find(uuid.UUID) (*user.Role, error) {
	return repo.getMockRole()
}

func (repo *mockRoleRepo) getMockRole() (*user.Role, error) {
	roleID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	role := &user.Role{
		roleID,
		"tester",
		[]string{"create-user"},
	}
	return role, nil
}

func (repo *mockRoleRepo) Delete(uuid.UUID) error {
	return nil
}

func (repo *mockRoleRepo) FindByUser(userID uuid.UUID) ([]*user.Role, error) {
	role, err := repo.getMockRole()
	if err != nil {
		return nil, err
	}
	return []*user.Role{role}, nil
}

func TestUserHasPermission(t *testing.T) {
	uc := &authorizationUseCase{&mockRoleRepo{}}

	userID, err := uuid.NewV4()
	if err != nil {
		t.Error(err)
	}

	permitted, err := uc.IsUserAuthorized(userID, "create-user")
	if err != nil {
		t.Error(err)
	}

	if !permitted {
		t.Errorf("Should be permitted to create-user")
	}

	permitted, err = uc.IsUserAuthorized(userID, "update-user")
	if err != nil {
		t.Error(err)
	}

	if permitted {
		t.Errorf("Should not be permitted to update-user")
	}
}
