package service

import (
	"testing"
	"time"

	"backend/core/entity"
	"backend/pkg/errs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPatientRepository is a mock of PatientRepository interface
type MockPatientRepository struct {
	mock.Mock
}

func (m *MockPatientRepository) CreatePatient(patient *entity.Patient) error {
	args := m.Called(patient)
	return args.Error(0)
}

func (m *MockPatientRepository) GetPatientById(id string) (*entity.Patient, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Patient), args.Error(1)
}

func (m *MockPatientRepository) GetPatients(hospitalId, firstname, lastname, middlename, nationId, passportId, dateOfBirth, phone, email string, page, limit int) ([]entity.Patient, error) {
	args := m.Called(hospitalId, firstname, lastname, middlename, nationId, passportId, dateOfBirth, phone, email, page, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Patient), args.Error(1)
}

// Test RegisterPatient - Success
func TestPatientService_RegisterPatient_Success(t *testing.T) {
	mockRepo := new(MockPatientRepository)
	
	patient := &entity.Patient{
		HospitalId:   "HOSP0001",
		FirstNameTh:  "นายเอ",
		LastNameTh:   "สมชาย",
		FirstNameEn:  "A",
		LastNameEn:   "Somchai",
		NationalId:   "1111111111111",
		Birthday:     time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	
	mockRepo.On("CreatePatient", patient).Return(nil)
	
	service := NewPatientService(mockRepo)
	err := service.RegisterPatient(patient)
	
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// Test RegisterPatient - Database Error
func TestPatientService_RegisterPatient_DatabaseError(t *testing.T) {
	mockRepo := new(MockPatientRepository)
	
	patient := &entity.Patient{
		HospitalId:  "HOSP0001",
		FirstNameTh: "นายเอ",
	}
	
	mockRepo.On("CreatePatient", patient).Return(errs.ErrDatabaseError)
	
	service := NewPatientService(mockRepo)
	err := service.RegisterPatient(patient)
	
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

// Test GetPatientInfo - Success
func TestPatientService_GetPatientInfo_Success(t *testing.T) {
	mockRepo := new(MockPatientRepository)
	
	patient := &entity.Patient{
		Id:          "PAT001",
		HospitalId:  "HOSP0001",
		FirstNameTh: "นายเอ",
		LastNameTh:  "สมชาย",
		FirstNameEn: "A",
		LastNameEn:  "Somchai",
		NationalId:  "1111111111111",
	}
	
	mockRepo.On("GetPatientById", "PAT001").Return(patient, nil)
	
	service := NewPatientService(mockRepo)
	result, err := service.GetPatientInfo("PAT001")
	
	assert.NoError(t, err)
	assert.Equal(t, "PAT001", result.Id)
	assert.Equal(t, "นายเอ", result.FirstNameTh)
	mockRepo.AssertExpectations(t)
}

// Test GetPatientInfo - Not Found
func TestPatientService_GetPatientInfo_NotFound(t *testing.T) {
	mockRepo := new(MockPatientRepository)
	mockRepo.On("GetPatientById", "INVALID").Return(nil, errs.ErrNotFound)
	
	service := NewPatientService(mockRepo)
	result, err := service.GetPatientInfo("INVALID")
	
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// Test GetPatients - Success
func TestPatientService_GetPatients_Success(t *testing.T) {
	mockRepo := new(MockPatientRepository)
	
	patients := []entity.Patient{
		{
			Id:         "PAT001",
			HospitalId: "HOSP0001",
			FirstNameTh: "นายเอ",
			FirstNameEn: "A",
			NationalId: "1111111111111",
		},
		{
			Id:         "PAT002",
			HospitalId: "HOSP0001",
			FirstNameTh: "นางบี",
			FirstNameEn: "B",
			NationalId: "2222222222222",
		},
	}
	
	mockRepo.On("GetPatients", "HOSP0001", "", "", "", "", "", "", "", "", 1, 10).Return(patients, nil)
	
	service := NewPatientService(mockRepo)
	result, err := service.GetPatients("HOSP0001", "", "", "", "", "", "", "", "", 1, 10)
	
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	mockRepo.AssertExpectations(t)
}

// Test GetPatients - With Firstname Filter
func TestPatientService_GetPatients_WithFirstnameFilter(t *testing.T) {
	mockRepo := new(MockPatientRepository)
	
	patients := []entity.Patient{
		{
			Id:          "PAT001",
			HospitalId:  "HOSP0001",
			FirstNameTh: "นายเอ",
			FirstNameEn: "A",
			NationalId:  "1111111111111",
		},
	}
	
	mockRepo.On("GetPatients", "HOSP0001", "นายเอ", "", "", "", "", "", "", "", 1, 10).Return(patients, nil)
	
	service := NewPatientService(mockRepo)
	result, err := service.GetPatients("HOSP0001", "นายเอ", "", "", "", "", "", "", "", 1, 10)
	
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result))
	mockRepo.AssertExpectations(t)
}

// Test GetPatients - Empty Result
func TestPatientService_GetPatients_EmptyResult(t *testing.T) {
	mockRepo := new(MockPatientRepository)
	
	patients := []entity.Patient{}
	mockRepo.On("GetPatients", "HOSP9999", "", "", "", "", "", "", "", "", 1, 10).Return(patients, nil)
	
	service := NewPatientService(mockRepo)
	result, err := service.GetPatients("HOSP9999", "", "", "", "", "", "", "", "", 1, 10)
	
	assert.NoError(t, err)
	assert.Equal(t, 0, len(result))
	mockRepo.AssertExpectations(t)
}

// Test GetPatients - Database Error
func TestPatientService_GetPatients_DatabaseError(t *testing.T) {
	mockRepo := new(MockPatientRepository)
	mockRepo.On("GetPatients", "HOSP0001", "", "", "", "", "", "", "", "", 1, 10).Return(nil, errs.ErrDatabaseError)
	
	service := NewPatientService(mockRepo)
	result, err := service.GetPatients("HOSP0001", "", "", "", "", "", "", "", "", 1, 10)
	
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// Test GetPatients - With Pagination
func TestPatientService_GetPatients_WithPagination(t *testing.T) {
	mockRepo := new(MockPatientRepository)
	
	patients := []entity.Patient{
		{
			Id:          "PAT003",
			HospitalId:  "HOSP0001",
			FirstNameTh: "นายซี",
			FirstNameEn: "C",
			NationalId:  "3333333333333",
		},
	}
	
	mockRepo.On("GetPatients", "HOSP0001", "", "", "", "", "", "", "", "", 2, 1).Return(patients, nil)
	
	service := NewPatientService(mockRepo)
	result, err := service.GetPatients("HOSP0001", "", "", "", "", "", "", "", "", 2, 1)
	
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result))
	mockRepo.AssertExpectations(t)
}
