package fico

import uuid "github.com/satori/go.uuid"

type Chart struct {
	ID     uuid.UUID
	Name   string
	Active bool
}
