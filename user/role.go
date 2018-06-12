package user

import uuid "github.com/satori/go.uuid"

type Role struct {
	ID          uuid.UUID
	RoleName    string
	Permissions []string
}

func (r *Role) AddPermission(perm string) {
	r.Permissions = append(r.Permissions, perm)
}

func (r *Role) HasPermission(perm string) bool {
	for _, item := range r.Permissions {
		if item == perm {
			return true
		}
	}
	return false
}

type RoleUser struct {
	UserID uuid.UUID
	RoleID uuid.UUID
}

type RoleRepo interface {
	Create(*Role) error
	Update(*Role) error
	Find(uuid.UUID) (*Role, error)
	Delete(uuid.UUID) error

	FindByUser(userID uuid.UUID) ([]*Role, error)
}
