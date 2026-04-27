package service

import (
	"testing"

	"backend/core/entity"
	"backend/pkg/errs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHospitalRepository is a mock of HospitalRepository interface
type MockHospitalRepository struct {
	mock.Mock
}

func (m *MockHospitalRepository) CreateHospital(hospital *entity.Hospital) (*entity.Hospital, error) {
	args := m.Called(hospital)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Hospital), args.Error(1)
}

func (m *MockHospitalRepository) GetHospitalById(id string) (*entity.Hospital, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Hospital), args.Error(1)
}

func (m *MockHospitalRepository) GetAllHospitals() ([]entity.Hospital, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Hospital), args.Error(1)
}

// Test RegisterHospital - Success
func TestHospitalService_RegisterHospital_Success(t *testing.T) {
	mockRepo := new(MockHospitalRepository)

	hospital := &entity.Hospital{
		Name: "Hospital A",
	}

	resultHospital := &entity.Hospital{
		Id:   "HOSP0001",
		Name: "Hospital A",
	}

	mockRepo.On("CreateHospital", hospital).Return(resultHospital, nil)

	service := NewHospitalService(mockRepo)
	result, err := service.RegisterHospital(hospital)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "HOSP0001", result.Id)
	mockRepo.AssertExpectations(t)
}

// Test RegisterHospital - Database Error
func TestHospitalService_RegisterHospital_DatabaseError(t *testing.T) {
	mockRepo := new(MockHospitalRepository)

	hospital := &entity.Hospital{
		Name: "Hospital A",
	}

	mockRepo.On("CreateHospital", hospital).Return(nil, errs.ErrDatabaseError)

	service := NewHospitalService(mockRepo)
	result, err := service.RegisterHospital(hospital)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// Test GetHospitalInfo - Success
func TestHospitalService_GetHospitalInfo_Success(t *testing.T) {
	mockRepo := new(MockHospitalRepository)

	hospital := &entity.Hospital{
		Id:   "HOSP0001",
		Name: "Hospital A",
	}

	mockRepo.On("GetHospitalById", "HOSP0001").Return(hospital, nil)

	service := NewHospitalService(mockRepo)
	result, err := service.GetHospitalInfo("HOSP0001")

	assert.NoError(t, err)
	assert.Equal(t, "HOSP0001", result.Id)
	assert.Equal(t, "Hospital A", result.Name)
	mockRepo.AssertExpectations(t)
}

// Test GetHospitalInfo - Not Found
func TestHospitalService_GetHospitalInfo_NotFound(t *testing.T) {
	mockRepo := new(MockHospitalRepository)
	mockRepo.On("GetHospitalById", "INVALID").Return(nil, errs.ErrHospitalNotFound)

	service := NewHospitalService(mockRepo)
	result, err := service.GetHospitalInfo("INVALID")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// Test GetAllHospitals - Success
func TestHospitalService_GetAllHospitals_Success(t *testing.T) {
	mockRepo := new(MockHospitalRepository)

	hospitals := []entity.Hospital{
		{Id: "HOSP0001", Name: "Hospital A"},
		{Id: "HOSP0002", Name: "Hospital B"},
	}

	mockRepo.On("GetAllHospitals").Return(hospitals, nil)

	service := NewHospitalService(mockRepo)
	result, err := service.GetAllHospitals()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	mockRepo.AssertExpectations(t)
}

// Test GetAllHospitals - Empty List
func TestHospitalService_GetAllHospitals_EmptyList(t *testing.T) {
	mockRepo := new(MockHospitalRepository)

	hospitals := []entity.Hospital{}
	mockRepo.On("GetAllHospitals").Return(hospitals, nil)

	service := NewHospitalService(mockRepo)
	result, err := service.GetAllHospitals()

	assert.NoError(t, err)
	assert.Equal(t, 0, len(result))
	mockRepo.AssertExpectations(t)
}

// Test GetAllHospitals - Database Error
func TestHospitalService_GetAllHospitals_DatabaseError(t *testing.T) {
	mockRepo := new(MockHospitalRepository)
	mockRepo.On("GetAllHospitals").Return(nil, errs.ErrDatabaseError)

	service := NewHospitalService(mockRepo)
	result, err := service.GetAllHospitals()

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
