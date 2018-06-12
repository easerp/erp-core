package sd

import (
	"github.com/rhymond/go-money"
	"github.com/satori/go.uuid"
	"time"
)

type SO struct {
	ID             uuid.UUID
	DocumentNumber string
	CustomerID     uuid.UUID
	Date           time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	IsApplied      bool
	IsVoid         bool
}

type SODetail struct {
	ID         uuid.UUID
	LineNumber int
	ProductID  uuid.UUID
	Qty        int64
	UnitPrice  *money.Money
	Discount   *money.Money
}

func (sod *SODetail) GetSubTotal() (*money.Money, error) {
	return sod.UnitPrice.Multiply(sod.Qty).Subtract(sod.Discount)
}

type SORepo interface {
	Create(*SO) error
	Update(*SO) error
	Delete(uuid.UUID) error
	Find(uuid.UUID) (*SO, error)
}
