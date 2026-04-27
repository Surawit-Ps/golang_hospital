package ports

import (
	"backend/core/entity"
)

type HospitalRepository interface {
	CreateHospital(hospital *entity.Hospital) (*entity.Hospital, error)
	GetHospitalById(id string) (*entity.Hospital, error)
	GetAllHospitals() ([]entity.Hospital, error)
}

type HospitalService interface {
	RegisterHospital(hospital *entity.Hospital) (*entity.Hospital, error)
	GetHospitalInfo(id string) (*entity.Hospital, error)
	GetAllHospitals() ([]entity.Hospital, error)
}