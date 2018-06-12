package inv

import (
	"github.com/satori/go.uuid"
	"time"
)

type GR struct {
	ID             uuid.UUID
	DocumentNumber string
	OrderNumber    string
	OrderType      string
	SupplierID     uuid.UUID
	Date           time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	IsApplied      bool
	IsVoid         bool

	Details []*GRDetail
}

type GRDetail struct {
	ID          uuid.UUID
	LineNumber  int
	ProductID   uuid.UUID
	BatchNo     string
	SerialNo    string
	ExpiredDate time.Time
	Qty         int
}

type GRRepo interface {
	Create(*GR) error
	Update(*GR) error
	Delete(uuid.UUID) error
	Find(uuid.UUID) (*GR, error)
}
