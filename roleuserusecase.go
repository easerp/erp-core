package core

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type RoleUserUseCase interface {
	AddRoleToUser(userID, roleID uuid.UUID) error
	RemoveRoleFromUser(userID, roleID uuid.UUID) error
}

type roleUserUseCase struct {
	repo     RoleUserRepo
	userRepo UserRepo
	roleRepo RoleRepo
}

func (uc *roleUserUseCase) AddRoleToUser(userID, roleID uuid.UUID) (err error) {

	_, err = uc.userRepo.Find(userID)
	if err != nil {
		return
	}

	_, err = uc.roleRepo.Find(roleID)
	if err != nil {
		return
	}

	hasRole, err := uc.repo.UserHasRole(userID, roleID)
	if err != nil {
		return
	}

	if hasRole {
		err = errors.New("User already has this role")
		return
	}

	err = uc.repo.Attach(userID, roleID)
	return
}

func (uc *roleUserUseCase) RemoveRoleFromUser(userID, roleID uuid.UUID) (err error) {
	_, err = uc.userRepo.Find(userID)
	if err != nil {
		return
	}

	_, err = uc.roleRepo.Find(roleID)
	if err != nil {
		return
	}

	hasRole, err := uc.repo.UserHasRole(userID, roleID)
	if err != nil {
		return
	}

	if !hasRole {
		err = errors.New("User does not have this role")
		return
	}

	err = uc.repo.Detach(userID, roleID)
	return
}
