package usecases

import (
	"github.com/flexrp/flexrp.core/sd"
	"github.com/satori/go.uuid"
	"errors"
)

type SOUseCase interface {
	FindSO(uuid.UUID) (*sd.SO, error)
	CreateSO(*sd.SO) error
	UpdateSO(*sd.SO) error
	ApplySO(soID uuid.UUID) error
	VoidSO(soID uuid.UUID) error
}

type soUseCase struct {
	repo sd.SORepo
}

func (uc *soUseCase) FindSO(soID uuid.UUID) (so *sd.SO, err error) {
	return uc.repo.Find(soID)
}

func (uc *soUseCase) CreateSO(so *sd.SO) (err error) {
	err = uc.repo.Create(so)
	return
}

func (uc *soUseCase) UpdateSO(so *sd.SO) (err error) {
	err = uc.repo.Update(so)
	return
}

func (uc *soUseCase) ApplySO(soID uuid.UUID) (err error) {
	so, err := uc.repo.Find(soID)
	if err != nil {
		return
	}
	if so.IsApplied {
		err = errors.New("Document already applied")
		return
	}
	if so.IsVoid {
		err = errors.New("Document already void")
		return
	}
	so.IsApplied = true
	err = uc.repo.Update(so)
	/**
	do apply action hook
	*/
	return
}

func (uc *soUseCase) VoidSO(soID uuid.UUID) (err error) {
	so, err := uc.repo.Find(soID)
	if err != nil {
		return
	}
	if so.IsVoid {
		err = errors.New("Document already void")
		return
	}
	so.IsVoid = true
	err = uc.repo.Update(so)
	/**
	do void action hook
	*/
	return
}
