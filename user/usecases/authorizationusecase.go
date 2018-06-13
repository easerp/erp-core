package usecases

import (
	"errors"

	"github.com/flexrp/flexrp.core/user"
	uuid "github.com/satori/go.uuid"
)

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
		//TODO
		//unauthorized error type
		err = errors.New("Unauthorized")
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
