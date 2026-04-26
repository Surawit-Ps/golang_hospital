package service


import (	"backend/core/entity"
	"backend/core/ports"
)

type HospitalServiceImpl struct {
	hospitalPort ports.HospitalPort
}

func NewHospitalService(hospitalPort ports.HospitalPort) ports.HospitalService {
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

func (s *HospitalServiceImpl) GetHospitalByCode(code string) (*entity.Hospital, error) {
	return s.hospitalPort.GetHospitalByCode(code)
}
