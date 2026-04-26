package service

import (
	"backend/core/entity"
	"backend/core/ports"
)

type StaffServiceImpl struct {
	staffPort ports.StaffPort
}

func NewStaffService(staffPort ports.StaffPort) ports.StaffService {
	return &StaffServiceImpl{
		staffPort: staffPort,
	}
}

func (s *StaffServiceImpl) RegisterStaff(staff *entity.Staff) (*entity.Staff, error) {
	return s.staffPort.CreateStaff(staff)
}

func (s *StaffServiceImpl) LoginStaff(username, password string) (*entity.Staff, error) {
	staff, err := s.staffPort.GetStaffByUsername(username)
	if err != nil {
		return nil, err
	}
	if staff.Password != password {
		return nil, err
	}
	return staff, nil
}