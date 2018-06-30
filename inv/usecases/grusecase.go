package usecases

import (
	"errors"

	"github.com/easerp/erp-core/inv"
	"github.com/satori/go.uuid"
)

type GRUseCase interface {
	FindGR(grID uuid.UUID) (*inv.GR, error)
	CreateGR(*inv.GR) error
	UpdateGR(*inv.GR) error
	ApplyGR(grID uuid.UUID) error
	VoidGR(grID uuid.UUID) error
}

type grUseCase struct {
	repo inv.GRRepo
}

func (uc *grUseCase) FindGR(grID uuid.UUID) (gr *inv.GR, err error) {
	gr, err = uc.repo.Find(grID)
	return
}

func (uc *grUseCase) CreateGR(gr *inv.GR) (err error) {
	err = uc.repo.Create(gr)
	return
}

func (uc *grUseCase) UpdateGR(gr *inv.GR) (err error) {
	err = uc.repo.Update(gr)
	return
}

func (uc *grUseCase) ApplyGR(grID uuid.UUID) (err error) {
	gr, err := uc.repo.Find(grID)
	if err != nil {
		return
	}
	if gr.IsApplied {
		err = errors.New("Document already applied")
		return
	}
	if gr.IsVoid {
		err = errors.New("Document already void")
		return
	}
	gr.IsApplied = true
	err = uc.repo.Update(gr)
	/**
	do apply action hook
	*/
	return
}

func (uc *grUseCase) VoidGR(grID uuid.UUID) (err error) {
	gr, err := uc.repo.Find(grID)
	if err != nil {
		return
	}
	if gr.IsVoid {
		err = errors.New("Document already void")
		return
	}
	gr.IsVoid = true
	err = uc.repo.Update(gr)
	/**
	do void action hook
	*/
	return
}
