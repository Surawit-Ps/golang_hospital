package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/core/entity"
	"backend/pkg/errs"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	h "backend/adapter/handler"
)

// MockHospitalService is a mock of HospitalService interface
type MockHospitalService struct {
	mock.Mock
}

func (m *MockHospitalService) RegisterHospital(hospital *entity.Hospital) (*entity.Hospital, error) {
	args := m.Called(hospital)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Hospital), args.Error(1)
}

func (m *MockHospitalService) GetHospitalInfo(id string) (*entity.Hospital, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Hospital), args.Error(1)
}

func (m *MockHospitalService) GetAllHospitals() ([]entity.Hospital, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Hospital), args.Error(1)
}

// Test RegisterHospital - Success
func TestRegisterHospital_Success(t *testing.T) {
	mockService := new(MockHospitalService)

	newHospital := &entity.Hospital{
		Id:   "HOSP0001",
		Name: "Hospital Test",
	}

	mockService.On("RegisterHospital", newHospital).Return(newHospital, nil)

	handler := h.NewHospitalHandler(mockService)

	body, _ := json.Marshal(newHospital)
	req := httptest.NewRequest(http.MethodPost, "/hospitals", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Create Gin context
	c := createTestContext(w, req)
	handler.RegisterHospital(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

// Test RegisterHospital - Invalid JSON
func TestRegisterHospital_InvalidJSON(t *testing.T) {
	mockService := new(MockHospitalService)
	handler := h.NewHospitalHandler(mockService)

	req := httptest.NewRequest(http.MethodPost, "/hospitals", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c := createTestContext(w, req)
	handler.RegisterHospital(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test GetHospitalInfo - Success
func TestGetHospitalInfo_Success(t *testing.T) {
	mockService := new(MockHospitalService)

	hospital := &entity.Hospital{
		Id:   "HOSP0001",
		Name: "Hospital A",
	}

	mockService.On("GetHospitalInfo", "HOSP0001").Return(hospital, nil)

	handler := h.NewHospitalHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/hospitals/HOSP0001", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "HOSP0001"})
	handler.GetHospitalInfo(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

// Test GetHospitalInfo - Not Found
func TestGetHospitalInfo_NotFound(t *testing.T) {
	mockService := new(MockHospitalService)
	mockService.On("GetHospitalInfo", "INVALID").Return(nil, errs.ErrHospitalNotFound)

	handler := h.NewHospitalHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/hospitals/INVALID", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "INVALID"})
	handler.GetHospitalInfo(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}

// Test GetAllHospitals - Success
func TestGetAllHospitals_Success(t *testing.T) {
	mockService := new(MockHospitalService)

	hospitals := []entity.Hospital{
		{Id: "HOSP0001", Name: "Hospital A"},
		{Id: "HOSP0002", Name: "Hospital B"},
	}

	mockService.On("GetAllHospitals").Return(hospitals, nil)

	handler := h.NewHospitalHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/hospitals", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	handler.GetAllHospitals(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

// Test GetAllHospitals - Empty List
func TestGetAllHospitals_EmptyList(t *testing.T) {
	mockService := new(MockHospitalService)

	hospitals := []entity.Hospital{}
	mockService.On("GetAllHospitals").Return(hospitals, nil)

	handler := h.NewHospitalHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/hospitals", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	handler.GetAllHospitals(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

// Test GetAllHospitals - Database Error
func TestGetAllHospitals_DatabaseError(t *testing.T) {
	mockService := new(MockHospitalService)
	mockService.On("GetAllHospitals").Return(nil, errs.ErrDatabaseError)

	handler := h.NewHospitalHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/hospitals", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	handler.GetAllHospitals(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
}
