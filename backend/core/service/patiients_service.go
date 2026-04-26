package service

import (
	"backend/core/entity"
	"backend/core/ports"
)

type PatientServiceImpl struct {
	patientPort ports.PatientPort
}

func NewPatientService(patientPort ports.PatientPort) ports.PatientService {
	return &PatientServiceImpl{
		patientPort: patientPort,
	}
}

func (s *PatientServiceImpl) RegisterPatient(patient *entity.Patient) (*entity.Patient, error) {
	return s.patientPort.CreatePatient(patient)
}

func (s *PatientServiceImpl) GetPatientInfo(id string) (*entity.Patient, error) {
	return s.patientPort.GetPatientById(id)
}

func (s *PatientServiceImpl) GetPatientsByHospital(hospitalId string) ([]*entity.Patient, error) {
	return s.patientPort.GetPatientsByHospitalId(hospitalId)
}