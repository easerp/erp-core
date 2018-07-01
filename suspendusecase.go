package core

import (
	uuid "github.com/satori/go.uuid"
)

type SuspendUseCase interface {
	SuspendUser(userID uuid.UUID) error
	UnsuspendUser(userID uuid.UUID) error
}

type suspendUseCase struct {
	repo UserRepo
}

func (uc *suspendUseCase) SuspendUser(userID uuid.UUID) (err error) {
	usr, err := uc.repo.Find(userID)
	if err != nil {
		return
	}

	usr.IsSuspended = true
	err = uc.repo.Update(usr)
	return
}

func (uc *suspendUseCase) UnsuspendUser(userID uuid.UUID) (err error) {
	usr, err := uc.repo.Find(userID)
	if err != nil {
		return
	}

	usr.IsSuspended = false
	err = uc.repo.Update(usr)
	return
}
