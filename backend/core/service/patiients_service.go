package service

import (
	"backend/core/entity"
	"backend/core/ports"
)

type PatientServiceImpl struct {
	patientPort ports.PatientRepository	
}

func NewPatientService(patientPort ports.PatientRepository) *PatientServiceImpl {
	return &PatientServiceImpl{
		patientPort: patientPort,
	}
}
func (s *PatientServiceImpl) RegisterPatient(patient *entity.Patient) error {
	return s.patientPort.CreatePatient(patient)
}

func (s *PatientServiceImpl) GetPatientInfo(id string) (*entity.Patient, error) {
	return s.patientPort.GetPatientById(id)
}

func (s *PatientServiceImpl) GetPatients(hospitalId string, firstname string, lastname string, middlename string, nationId string, passportId string,dateOfBirth string,phone string,email string, page int, limit int) ([]entity.Patient, error) {
	return s.patientPort.GetPatients(hospitalId, firstname, lastname, middlename, nationId, passportId,dateOfBirth,phone,email, page, limit)
}