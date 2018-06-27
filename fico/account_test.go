package fico

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestNewAccount(t *testing.T) {
	coaID, err := uuid.NewV4()
	if err != nil {
		t.Error(err)
	}

	accParent, err := NewAccount(
		"1000",
		"Asset",
		coaID,
		uuid.Nil,
	)

	if err != nil {
		t.Error(err)
	}

	acc, err := NewAccount(
		"1100",
		"Liquid Asset",
		coaID,
		accParent.ID,
	)

	if err != nil {
		t.Error(err)
	}

	if acc.AccountNo != "1100" {
		t.Error("Mismatch AccountNo")
	}

	if acc.Name != "Liquid Asset" {
		t.Error("Mismatch Name")
	}

	if acc.ChartID != coaID {
		t.Error("Mismatch ChartID")
	}

	if acc.ParentID != accParent.ID {
		t.Error("Mismatch Parent")
	}

}
