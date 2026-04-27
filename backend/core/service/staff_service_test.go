package service

import (
	"testing"

	"backend/core/entity"
	"backend/core/middleware"
	"backend/pkg/errs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStaffRepository is a mock of StaffRepository interface
type MockStaffRepository struct {
	mock.Mock
}

func (m *MockStaffRepository) CreateStaff(staff *entity.Staff) error {
	args := m.Called(staff)
	return args.Error(0)
}

func (m *MockStaffRepository) GetStaffById(id string) (*entity.Staff, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Staff), args.Error(1)
}

func (m *MockStaffRepository) GetStaffByUsernameAndHospitalId(username, hospitalId string) (*entity.Staff, error) {
	args := m.Called(username, hospitalId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Staff), args.Error(1)
}

func (m *MockStaffRepository) GetUser() (*entity.Staff, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Staff), args.Error(1)
}
// Test RegisterStaff - Success
func TestStaffService_RegisterStaff_Success(t *testing.T) {
	mockRepo := new(MockStaffRepository)

	staff := &entity.Staff{
		Username:   "doctor1",
		Password:   "hashed_password",
		HospitalId: "HOSP0001",
	}

	mockRepo.On("CreateStaff", staff).Return(nil)

	service := NewStaffService(mockRepo)
	err := service.RegisterStaff(staff)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// Test RegisterStaff - Database Error
func TestStaffService_RegisterStaff_DatabaseError(t *testing.T) {
	mockRepo := new(MockStaffRepository)

	staff := &entity.Staff{
		Username:   "doctor1",
		Password:   "hashed_password",
		HospitalId: "HOSP0001",
	}

	mockRepo.On("CreateStaff", staff).Return(errs.ErrDatabaseError)

	service := NewStaffService(mockRepo)
	err := service.RegisterStaff(staff)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

// Test LoginStaff - Success
func TestStaffService_LoginStaff_Success(t *testing.T) {
	mockRepo := new(MockStaffRepository)

	hashedPassword, _ := middleware.HashPassword("123456")

	staff := &entity.Staff{
		Id:         "STAFF001",
		Username:   "staff1",
		Password:   hashedPassword,
		HospitalId: "HOSP0001",
		Role:       "staff",
	}

	mockRepo.On("GetStaffByUsernameAndHospitalId", "staff1", "HOSP0001").Return(staff, nil)

	service := NewStaffService(mockRepo)
	token, loginStaff, err := service.LoginStaff("staff1", "123456", "HOSP0001")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, "STAFF001", loginStaff.Id)
	mockRepo.AssertExpectations(t)
}

// Test LoginStaff - Invalid Credentials
func TestStaffService_LoginStaff_InvalidCredentials(t *testing.T) {
	mockRepo := new(MockStaffRepository)

	hashedPassword, _ := middleware.HashPassword("123456")

	staff := &entity.Staff{
		Id:         "STAFF001",
		Username:   "staff1",
		Password:   hashedPassword,
		HospitalId: "HOSP0001",
	}

	mockRepo.On("GetStaffByUsernameAndHospitalId", "staff1", "HOSP0001").Return(staff, nil)

	service := NewStaffService(mockRepo)
	token, loginStaff, err := service.LoginStaff("staff1", "wrongpassword", "HOSP0001")

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Nil(t, loginStaff)
	mockRepo.AssertExpectations(t)
}

// Test LoginStaff - User Not Found
func TestStaffService_LoginStaff_UserNotFound(t *testing.T) {
	mockRepo := new(MockStaffRepository)
	mockRepo.On("GetStaffByUsernameAndHospitalId", "invalid", "HOSP0001").Return(nil, errs.ErrNotFound)

	service := NewStaffService(mockRepo)
	token, loginStaff, err := service.LoginStaff("invalid", "123456", "HOSP0001")

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Nil(t, loginStaff)
	mockRepo.AssertExpectations(t)
}

// Test GetUser - Success
func TestStaffService_GetUser_Success(t *testing.T) {
	mockRepo := new(MockStaffRepository)

	staff := &entity.Staff{
		Id:         "STAFF001",
		Username:   "staff1",
		HospitalId: "HOSP0001",
	}

	mockRepo.On("GetStaffById", "STAFF001").Return(staff, nil)

	service := NewStaffService(mockRepo)
	result, err := service.GetUser("STAFF001")

	assert.NoError(t, err)
	assert.Equal(t, "STAFF001", result.Id)
	mockRepo.AssertExpectations(t)
}

// Test GetUser - Not Found
func TestStaffService_GetUser_NotFound(t *testing.T) {
	mockRepo := new(MockStaffRepository)
	mockRepo.On("GetStaffById", "INVALID").Return(nil, errs.ErrNotFound)

	service := NewStaffService(mockRepo)
	result, err := service.GetUser("INVALID")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
