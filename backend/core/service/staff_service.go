package service

import (
	"backend/core/entity"
	"backend/core/middleware"
	"backend/core/ports"
	"backend/pkg/errs"
)

type staffService struct {
	staffPort ports.StaffRepository
}

func NewStaffService(staffPort ports.StaffRepository) *staffService {
	return &staffService{
		staffPort: staffPort,
	}
}

func (s *staffService) RegisterStaff(staff *entity.Staff) error {
	return s.staffPort.CreateStaff(staff)
}

func (s *staffService) LoginStaff(username, password, hospitalId string) (string, *entity.Staff, error) {
	staff, err := s.staffPort.GetStaffByUsernameAndHospitalId(username, hospitalId)
	if err != nil {
		return "", nil, err
	}
	ok := middleware.CheckPasswordHash([]byte(password), []byte(staff.Password))
	if !ok {
		return "", nil, errs.ErrInvalidCredentials
	}
	jwtWrapper := middleware.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "authService",
		ExpirationHours: 24,
	}
	token, err := jwtWrapper.GenerateToken(staff.Id, staff.Role, staff.HospitalId)
	if err != nil {
		return "", nil, errs.ErrInternalServerError
	}

	return token, staff, nil
}

func (s *staffService) GetUser(id string) (*entity.Staff, error) {
	staff, err := s.staffPort.GetStaffById(id)
	if err != nil {
		return nil, err
	}
	return staff, nil
}