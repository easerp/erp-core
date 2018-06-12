package usecases

import (
	"github.com/flexrp/flexrp.core/user"
	uuid "github.com/satori/go.uuid"
)

type PermissionUseCase interface {
	UserHasPermission(userID uuid.UUID, permission string) (bool, error)
}

type permissionUseCase struct {
	repo user.RoleRepo
}

func (uc *permissionUseCase) UserHasPermission(userID uuid.UUID, permission string) (
	permitted bool, err error) {

	permitted = false

	roles, err := uc.repo.FindByUser(userID)
	if err != nil {
		return
	}

	for _, role := range roles {
		if role.HasPermission(permission) {
			permitted = true
			return
		}
	}

	return
}
