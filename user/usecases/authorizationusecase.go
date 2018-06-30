package usecases

import (
	"github.com/easerp/erp-core/user"
	uuid "github.com/satori/go.uuid"
)

type AuthorizationError struct {
	err string
}

func (ae *AuthorizationError) Error() string {
	return ae.err
}

type AuthorizationUseCase interface {
	IsAuthorized(permission string) (bool, error)
	Authorize(permission string) error
}

type UserContext interface {
	GetUserID() (uuid.UUID, error)
}

type authorizationUseCase struct {
	repo user.RoleRepo
	ctx  UserContext
}

func (uc *authorizationUseCase) IsAuthorized(permission string) (
	permitted bool, err error) {

	userID, err := uc.ctx.GetUserID()
	if err != nil {
		return
	}

	permitted, err = uc.isUserAuthorized(userID, permission)
	return
}

func (uc *authorizationUseCase) Authorize(permission string) (err error) {
	authorized, err := uc.IsAuthorized(permission)
	if err != nil {
		return
	}

	if !authorized {
		return &AuthorizationError{"Unauthorized"}
		return
	}
	return
}

func (uc *authorizationUseCase) isUserAuthorized(userID uuid.UUID, permission string) (
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
