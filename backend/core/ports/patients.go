package ports

import (
	"backend/core/entity"
)

type PatientPort interface {
	CreatePatient(patient *entity.Patient) (*entity.Patient, error)
	GetPatientById(id string) (*entity.Patient, error)
	GetPatientsByHospitalId(hospitalId string) ([]*entity.Patient, error)
}

type PatientService interface {
	RegisterPatient(patient *entity.Patient) (*entity.Patient, error)
	GetPatientInfo(id string) (*entity.Patient, error)
	GetPatientsByHospital(hospitalId string) ([]*entity.Patient, error)
}