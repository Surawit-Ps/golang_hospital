package service


import (	"backend/core/entity"
	"backend/core/ports"
)

type HospitalServiceImpl struct {
	hospitalPort ports.HospitalRepository
}

func NewHospitalService(hospitalPort ports.HospitalRepository) *HospitalServiceImpl {
	return &HospitalServiceImpl{
		hospitalPort: hospitalPort,
	}
}

func (s *HospitalServiceImpl) RegisterHospital(hospital *entity.Hospital) (*entity.Hospital, error) {
	return s.hospitalPort.CreateHospital(hospital)
}

func (s *HospitalServiceImpl) GetHospitalInfo(id string) (*entity.Hospital, error) {
	return s.hospitalPort.GetHospitalById(id)
}

func (s *HospitalServiceImpl) GetAllHospitals() ([]entity.Hospital, error) {
	return s.hospitalPort.GetAllHospitals()
}