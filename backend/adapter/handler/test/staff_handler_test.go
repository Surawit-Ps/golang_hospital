package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/core/entity"
	"backend/pkg/errs"
	h "backend/adapter/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStaffService is a mock of StaffService interface
type MockStaffService struct {
	mock.Mock
}

func (m *MockStaffService) RegisterStaff(staff *entity.Staff) error {
	args := m.Called(staff)
	return args.Error(0)
}

func (m *MockStaffService) LoginStaff(username, password, hospitalId string) (string, *entity.Staff, error) {
	args := m.Called(username, password, hospitalId)
	return args.String(0), args.Get(1).(*entity.Staff), args.Error(2)
}

func (m *MockStaffService) GetUser(id string) (*entity.Staff, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Staff), args.Error(1)
}

// Test RegisterStaff - Success
func TestRegisterStaff_Success(t *testing.T) {
	mockService := new(MockStaffService)

	newStaff := &entity.Staff{
		Id:         "STAFF001",
		Username:   "doctor1",
		Password:   "hashed_password",
		HospitalId: "HOSP0001",
		Role:       "doctor",
	}

	mockService.On("RegisterStaff", newStaff).Return(nil)

	handler := h.NewStaffHandler(mockService)

	body, _ := json.Marshal(newStaff)
	req := httptest.NewRequest(http.MethodPost, "/staff", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c := createTestContext(w, req)
	// Set userID in context
	c.Set("userID", "STAFF001")

	handler.RegisterStaff(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

// Test RegisterStaff - Unauthorized
func TestRegisterStaff_Unauthorized(t *testing.T) {
	mockService := new(MockStaffService)
	handler := h.NewStaffHandler(mockService)

	newStaff := &entity.Staff{
		Username:   "doctor1",
		Password:   "hashed_password",
		HospitalId: "HOSP0001",
	}

	body, _ := json.Marshal(newStaff)
	req := httptest.NewRequest(http.MethodPost, "/staff", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c := createTestContext(w, req)
	// Don't set userID - should cause unauthorized error

	handler.RegisterStaff(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// Test RegisterStaff - Invalid JSON
func TestRegisterStaff_InvalidJSON(t *testing.T) {
	mockService := new(MockStaffService)
	handler := h.NewStaffHandler(mockService)

	req := httptest.NewRequest(http.MethodPost, "/staff", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c := createTestContext(w, req)
	c.Set("userID", "STAFF001")

	handler.RegisterStaff(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test LoginStaff - Success
func TestLoginStaff_Success(t *testing.T) {
	mockService := new(MockStaffService)

	loginData := map[string]string{
		"username":    "staff1",
		"password":    "123456",
		"hospital_id": "HOSP0001",
	}

	mockService.On("LoginStaff", "staff1", "123456", "HOSP0001").Return("token_string", &entity.Staff{Id: "STAFF001"}, nil)

	handler := h.NewStaffHandler(mockService)

	body, _ := json.Marshal(loginData)
	req := httptest.NewRequest(http.MethodPost, "/staff/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c := createTestContext(w, req)

	handler.LoginStaff(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

// Test LoginStaff - Invalid Credentials
func TestLoginStaff_InvalidCredentials(t *testing.T) {
	mockService := new(MockStaffService)

	loginData := map[string]string{
		"username":    "staff1",
		"password":    "wrongpassword",
		"hospital_id": "HOSP0001",
	}

	mockService.On("LoginStaff", "staff1", "wrongpassword", "HOSP0001").Return("", errs.ErrInvalidCredentials)

	handler := h.NewStaffHandler(mockService)

	body, _ := json.Marshal(loginData)
	req := httptest.NewRequest(http.MethodPost, "/staff/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c := createTestContext(w, req)

	handler.LoginStaff(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockService.AssertExpectations(t)
}

// Test GetUserLogin - Success
func TestGetUserLogin_Success(t *testing.T) {
	mockService := new(MockStaffService)

	staff := &entity.Staff{
		Id:         "STAFF001",
		Username:   "staff1",
		HospitalId: "HOSP0001",
		Role:       "staff",
	}

	mockService.On("GetUser", "STAFF001").Return(staff, nil)

	handler := h.NewStaffHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/staff", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	c.Set("userID", "STAFF001")

	handler.GetUserLogin(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

// Test GetUserLogin - Unauthorized
func TestGetUserLogin_Unauthorized(t *testing.T) {
	mockService := new(MockStaffService)
	handler := h.NewStaffHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/staff", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	// Don't set userID

	handler.GetUserLogin(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// Test GetUserLogin - Not Found
func TestGetUserLogin_NotFound(t *testing.T) {
	mockService := new(MockStaffService)
	mockService.On("GetUser", "INVALID").Return(nil, errs.ErrNotFound)

	handler := h.NewStaffHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/staff", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	c.Set("userID", "INVALID")

	handler.GetUserLogin(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}
