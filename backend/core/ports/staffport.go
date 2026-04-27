package ports

import (
	"backend/core/entity"
)

type StaffRepository interface {
	CreateStaff(staff *entity.Staff) error
	GetStaffById(id string) (*entity.Staff, error)
	GetStaffByUsernameAndHospitalId(username, hospitalId string) (*entity.Staff, error)
	GetUser() (*entity.Staff, error)
}

type StaffService interface {
	RegisterStaff(staff *entity.Staff) error
	LoginStaff(username, password, hospitalId string) (string, *entity.Staff, error)
	GetUser(id string) (*entity.Staff, error)
}