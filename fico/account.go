package fico

import uuid "github.com/satori/go.uuid"

type Account struct {
	ID        uuid.UUID
	AccountNo string
	Name      string
	ChartID   uuid.UUID
	ParentID  uuid.UUID
}

func NewAccount(accountNo, name string, chartID, parentID uuid.UUID) (*Account, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:        id,
		AccountNo: accountNo,
		Name:      name,
		ChartID:   chartID,
		ParentID:  parentID,
	}, nil
}
