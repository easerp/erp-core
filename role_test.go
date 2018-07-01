package core

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestHasPermission(t *testing.T) {
	roleID, err := uuid.NewV4()
	if err != nil {
		t.Error(err)
	}

	role := &Role{
		roleID,
		"tester",
		[]string{"create-user"},
	}

	if !role.HasPermission("create-user") {
		t.Errorf("Should have permission create-user")
	}

	if role.HasPermission("update-user") {
		t.Errorf("Should not have update-user")
	}

	role.AddPermission("update-user")

	if !role.HasPermission("update-user") {
		t.Errorf("Should have permission update-user")
	}
}
