package core

import uuid "github.com/satori/go.uuid"

type UserContext interface {
	GetUserID() (uuid.UUID, error)
}
