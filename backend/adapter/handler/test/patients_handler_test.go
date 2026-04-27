package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"backend/core/entity"
	"backend/pkg/errs"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	h "backend/adapter/handler"
)

// MockPatientService is a mock of PatientService interface
type MockPatientService struct {
	mock.Mock
}

func (m *MockPatientService) RegisterPatient(patient *entity.Patient) error {
	args := m.Called(patient)
	return args.Error(0)
}

func (m *MockPatientService) GetPatientInfo(id string) (*entity.Patient, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Patient), args.Error(1)
}

func (m *MockPatientService) GetPatients(hospitalId, firstname, lastname, middlename, nationId, passportId, dateOfBirth, phone, email string, page, limit int) ([]entity.Patient, error) {
	args := m.Called(hospitalId, firstname, lastname, middlename, nationId, passportId, dateOfBirth, phone, email, page, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Patient), args.Error(1)
}

// Test RegisterPatient - Success
func TestRegisterPatient_Success(t *testing.T) {
	mockService := new(MockPatientService)

	newPatient := &entity.Patient{
		HospitalId:  "HOSP0001",
		FirstNameTh: "นายเอ",
		LastNameTh:  "สมชาย",
		FirstNameEn: "A",
		LastNameEn:  "Somchai",
		NationalId:  "1111111111111",
		Birthday:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	mockService.On("RegisterPatient", newPatient).Return(nil)

	handler := h.NewPatientHandler(mockService)

	body, _ := json.Marshal(newPatient)
	req := httptest.NewRequest(http.MethodPost, "/patients", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c := createTestContext(w, req)
	c.Set("userID", "STAFF001")

	handler.RegisterPatient(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

// Test RegisterPatient - Unauthorized
func TestRegisterPatient_Unauthorized(t *testing.T) {
	mockService := new(MockPatientService)
	handler := h.NewPatientHandler(mockService)

	newPatient := &entity.Patient{
		HospitalId:  "HOSP0001",
		FirstNameTh: "นายเอ",
		LastNameTh:  "สมชาย",
	}

	body, _ := json.Marshal(newPatient)
	req := httptest.NewRequest(http.MethodPost, "/patients", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c := createTestContext(w, req)
	// Don't set userID

	handler.RegisterPatient(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// Test RegisterPatient - Invalid JSON
func TestRegisterPatient_InvalidJSON(t *testing.T) {
	mockService := new(MockPatientService)
	handler := h.NewPatientHandler(mockService)

	req := httptest.NewRequest(http.MethodPost, "/patients", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c := createTestContext(w, req)
	c.Set("userID", "STAFF001")

	handler.RegisterPatient(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test GetPatientInfo - Success
func TestGetPatientInfo_Success(t *testing.T) {
	mockService := new(MockPatientService)

	patient := &entity.Patient{
		Id:          "PAT001",
		HospitalId:  "HOSP0001",
		FirstNameTh: "นายเอ",
		LastNameTh:  "สมชาย",
		FirstNameEn: "A",
		LastNameEn:  "Somchai",
		NationalId:  "1111111111111",
	}

	mockService.On("GetPatientInfo", "PAT001").Return(patient, nil)

	handler := h.NewPatientHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/patients/PAT001", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	c.Set("userID", "STAFF001")
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "PAT001"})

	handler.GetPatientInfo(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

// Test GetPatientInfo - Not Found
func TestGetPatientInfo_NotFound(t *testing.T) {
	mockService := new(MockPatientService)
	mockService.On("GetPatientInfo", "INVALID").Return(nil, errs.ErrNotFound)

	handler := h.NewPatientHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/patients/INVALID", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	c.Set("userID", "STAFF001")
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "INVALID"})

	handler.GetPatientInfo(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}

// Test GetPatients - Success
func TestGetPatients_Success(t *testing.T) {
	mockService := new(MockPatientService)

	patients := []entity.Patient{
		{
			Id:          "PAT001",
			HospitalId:  "HOSP0001",
			FirstNameTh: "นายเอ",
			LastNameTh:  "สมชาย",
			FirstNameEn: "A",
			LastNameEn:  "Somchai",
			NationalId:  "1111111111111",
		},
		{
			Id:          "PAT002",
			HospitalId:  "HOSP0001",
			FirstNameTh: "นางบี",
			LastNameTh:  "สมหญิง",
			FirstNameEn: "B",
			LastNameEn:  "Somying",
			NationalId:  "2222222222222",
		},
	}

	mockService.On("GetPatients", "HOSP0001", "", "", "", "", "", "", "", "", 1, 10).Return(patients, nil)

	handler := h.NewPatientHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/patients?page=1&limit=10", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	c.Set("userID", "STAFF001")
	c.Set("HospitalID", "HOSP0001")

	handler.GetPatients(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

// Test GetPatients - Unauthorized
func TestGetPatients_Unauthorized(t *testing.T) {
	mockService := new(MockPatientService)
	handler := h.NewPatientHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/patients", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	// Don't set userID

	handler.GetPatients(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// Test GetPatients - Invalid Page Number
func TestGetPatients_InvalidPageNumber(t *testing.T) {
	mockService := new(MockPatientService)
	handler := h.NewPatientHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/patients?page=invalid&limit=10", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	c.Set("userID", "STAFF001")
	c.Set("HospitalID", "HOSP0001")

	handler.GetPatients(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test GetPatients - Invalid Limit Number
func TestGetPatients_InvalidLimitNumber(t *testing.T) {
	mockService := new(MockPatientService)
	handler := h.NewPatientHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/patients?page=1&limit=invalid", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	c.Set("userID", "STAFF001")
	c.Set("HospitalID", "HOSP0001")

	handler.GetPatients(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test GetPatients - Empty List
func TestGetPatients_EmptyList(t *testing.T) {
	mockService := new(MockPatientService)

	patients := []entity.Patient{}
	mockService.On("GetPatients", "HOSP0001", "", "", "", "", "", "", "", "", 1, 10).Return(patients, nil)

	handler := h.NewPatientHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/patients", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	c.Set("userID", "STAFF001")
	c.Set("HospitalID", "HOSP0001")

	handler.GetPatients(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

// Test GetPatients - With Filter
func TestGetPatients_WithFilter(t *testing.T) {
	mockService := new(MockPatientService)

	patients := []entity.Patient{
		{
			Id:          "PAT001",
			HospitalId:  "HOSP0001",
			FirstNameTh: "นายเอ",
			FirstNameEn: "A",
			NationalId:  "1111111111111",
		},
	}

	mockService.On("GetPatients", "HOSP0001", "นายเอ", "", "", "", "", "", "", "", 1, 10).Return(patients, nil)

	handler := h.NewPatientHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/patients?firstname=นายเอ&page=1&limit=10", nil)
	w := httptest.NewRecorder()

	c := createTestContext(w, req)
	c.Set("userID", "STAFF001")
	c.Set("HospitalID", "HOSP0001")

	handler.GetPatients(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
