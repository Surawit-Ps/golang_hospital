package ports

import (
	"backend/core/entity"
)

type StaffPort interface {
	CreateStaff(staff *entity.Staff) (*entity.Staff, error)
	GetStaffById(id string) (*entity.Staff, error)
	GetStaffByUsername(username string) (*entity.Staff, error)
}

type StaffService interface {
	RegisterStaff(staff *entity.Staff) (*entity.Staff, error)
	LoginStaff(username, password string) (*entity.Staff, error)
}