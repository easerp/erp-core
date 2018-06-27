package fico

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type GLPost struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	Remark    uuid.UUID
	Date      time.Time
}
