package ports

import (
	"backend/core/entity"
)

type PatientRepository interface {
	CreatePatient(patient *entity.Patient) error
	GetPatientById(id string) (*entity.Patient, error)
	GetPatients(hospitalId string, firstname string, lastname string, middlename string, nationId string, passportId string,dateOfBirth string,phone string,email string,page int, limit int) ([]entity.Patient, error)
}


type PatientService interface {
	RegisterPatient(patient *entity.Patient) error
	GetPatientInfo(id string) (*entity.Patient, error)
	GetPatients(hospitalId string, firstname string, lastname string, middlename string, nationId string, passportId string,dateOfBirth string,phone string,email string,page int, limit int) ([]entity.Patient, error)
}