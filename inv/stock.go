package inv

import (
	"github.com/satori/go.uuid"
	"time"
)

type Stock struct {
	ProductID uuid.UUID
	Qty       int
	Details   []*StockDetail
}

type StockDetail struct {
	ID          uuid.UUID
	SupplierID  uuid.UUID
	BatchNo     string
	SerialNo    string
	ExpiredDate time.Time
	Qty         int
}
