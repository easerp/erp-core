package sd

import (
	"testing"

	money "github.com/rhymond/go-money"
	uuid "github.com/satori/go.uuid"
)

func TestGetSubTotal(t *testing.T) {
	id, err := uuid.NewV4()
	if err != nil {
		t.Error(err)
	}

	productID, err := uuid.NewV4()
	if err != nil {
		t.Error(err)
	}

	sod := &SODetail{
		id,
		1,
		productID,
		10,
		money.New(1000, "IDR"),
		money.New(100, "IDR"),
	}

	expected := int64(9900)
	actual, err := sod.GetSubTotal()
	if err != nil {
		t.Error(err)
	}

	if expected != actual.Amount() {
		t.Errorf("Expected sub total: %d, actual: %d", expected, actual.Amount())
	}
}
