package sd

import (
	"time"

	"github.com/rhymond/go-money"
	"github.com/satori/go.uuid"
)

type SO struct {
	ID             uuid.UUID `validate:required`
	DocumentNumber string    `validate:required`
	CustomerID     uuid.UUID
	Date           time.Time `validate:required`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	IsApplied      bool
	IsVoid         bool
}

type SODetail struct {
	ID         uuid.UUID    `validate:required`
	LineNumber int          `validate:required`
	ProductID  uuid.UUID    `validate:required`
	Qty        int64        `validate:required`
	UnitPrice  *money.Money `validate:required`
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
