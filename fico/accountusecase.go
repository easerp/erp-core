package fico

import uuid "github.com/satori/go.uuid"

type AccountRepo interface {
	Find(uuid.UUID) (*Account, error)
	Create(*Account) error
	Update(*Account) error
	Delete(uuid.UUID) error
	FindByParentID(parentID uuid.UUID) ([]*Account, error)
	FindByChartID(chartID uuid.UUID) ([]*Account, error)
}

type AccountUseCase interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error

	GetRootAccounts() ([]*Account, error)
	GetAccount(uuid.UUID) (*Account, error)
	GetChildren(uuid.UUID) ([]*Account, error)

	AddChild(id uuid.UUID, child *Account) error
	DeleteAccount(id uuid.UUID) error
}

type accountUseCase struct {
	repo AccountRepo
}

func (uc *accountUseCase) GetRootAccounts() ([]*Account, error) {
	return uc.repo.FindByParentID(uuid.Nil)
}

func (uc *accountUseCase) GetAccount(id uuid.UUID) (acc *Account, err error) {
	acc, err = uc.repo.Find(id)
	if err != nil {
		return
	}
	return
}

func (uc *accountUseCase) GetChildren(parentID uuid.UUID) (children []*Account, err error) {
	children, err = uc.repo.FindByParentID(parentID)
	if err != nil {
		return
	}
	return
}

func (uc *accountUseCase) CreateAccount(acc *Account) error {
	err := uc.repo.Create(acc)
	if err != nil {
		return err
	}
	return nil
}

func (uc *accountUseCase) UpdateAccount(acc *Account) error {
	_, err := uc.repo.Find(acc.ID)
	if err != nil {
		return err
	}

	err = uc.repo.Update(acc)
	if err != nil {
		return err
	}
	return nil
}

func (uc *accountUseCase) AddChild(id uuid.UUID, child *Account) (err error) {
	parent, err := uc.GetAccount(id)
	if err != nil {
		return
	}

	child.ParentID = parent.ID

	err = uc.CreateAccount(child)
	return
}

func (uc *accountUseCase) DeleteAccount(id uuid.UUID) error {
	return uc.repo.Delete(id)
}
