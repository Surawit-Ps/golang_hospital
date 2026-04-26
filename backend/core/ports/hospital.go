package ports

import (
	"backend/core/entity"
)

type HospitalPort interface {
	CreateHospital(hospital *entity.Hospital) (*entity.Hospital, error)
	GetHospitalById(id string) (*entity.Hospital, error)
	GetHospitalByCode(code string) (*entity.Hospital, error)
}

type HospitalService interface {
	RegisterHospital(hospital *entity.Hospital) (*entity.Hospital, error)
	GetHospitalInfo(id string) (*entity.Hospital, error)
	GetHospitalByCode(code string) (*entity.Hospital, error)
}