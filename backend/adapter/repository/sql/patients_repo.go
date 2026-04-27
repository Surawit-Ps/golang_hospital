package repository

import (
	"backend/core/entity"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *patientRepository {
	return &patientRepository{db: db}
}

type Patient struct {
	Id           string `gorm:"primaryKey"`
	HospitalId   string `gorm:"index"`
	PatientHn    string `gorm:"uniqueIndex"`
	FirstNameTh  string
	LastNameTh   string
	MiddleNameTh string
	FirstNameEn  string
	MiddleNameEn string
	LastNameEn   string
	NationalId   string
	PassportId   string
	Birthday     time.Time
	Phone        string
	Email        string
	Gender       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func patientToEntity(patient *Patient) *entity.Patient {
	return &entity.Patient{
		Id:           patient.Id,
		HospitalId:   patient.HospitalId,
		PatientHn:    patient.PatientHn,
		FirstNameTh:  patient.FirstNameTh,
		LastNameTh:   patient.LastNameTh,
		MiddleNameTh: patient.MiddleNameTh,
		FirstNameEn:  patient.FirstNameEn,
		MiddleNameEn: patient.MiddleNameEn,
		LastNameEn:   patient.LastNameEn,
		NationalId:   patient.NationalId,
		PassportId:   patient.PassportId,
		Birthday:     patient.Birthday,
		Phone:        patient.Phone,
		Email:        patient.Email,
		Gender:       patient.Gender,
		CreatedAt:    patient.CreatedAt,
		UpdatedAt:    patient.UpdatedAt,
	}
}

func entityToPatient(patient *entity.Patient) *Patient {
	return &Patient{
		Id:           patient.Id,
		HospitalId:   patient.HospitalId,
		PatientHn:    patient.PatientHn,
		FirstNameTh:  patient.FirstNameTh,
		LastNameTh:   patient.LastNameTh,
		MiddleNameTh: patient.MiddleNameTh,
		FirstNameEn:  patient.FirstNameEn,
		MiddleNameEn: patient.MiddleNameEn,
		LastNameEn:   patient.LastNameEn,
		NationalId:   patient.NationalId,
		PassportId:   patient.PassportId,
		Birthday:     patient.Birthday,
		Phone:        patient.Phone,
		Email:        patient.Email,
		Gender:       patient.Gender,
		CreatedAt:    patient.CreatedAt,
		UpdatedAt:    patient.UpdatedAt,
	}
}

func (r *patientRepository) CreatePatient(patient *entity.Patient) error {
	patient.Id = uuid.New().String()
	patient.CreatedAt = time.Now()
	patient.UpdatedAt = time.Now()
	patient.PatientHn, _ = r.GeneratePatientCode()
	sqlPatient := entityToPatient(patient)
	if err := r.db.Create(sqlPatient).Error; err != nil {
		return err
	}
	return nil
}

func (r *patientRepository) GetPatientById(id string) (*entity.Patient, error) {
	var patient Patient
	if err := r.db.First(&patient, "id = ? OR hospital_id = ?", id, id).Error; err != nil {
		return nil, err
	}
	patientEntity := patientToEntity(&patient)
	return patientEntity, nil
}

func (r *patientRepository) GetPatients(hospitalId string, firstname string, lastname string, middlename string, nationId string, passportId string, dateOfBirth string, phone string, email string, page int, limit int) ([]entity.Patient, error) {
	var patients []Patient
	query := r.db.Where("hospital_id = ?", hospitalId)
	if firstname != "" {
		query = query.Where("first_name_th ILIKE ?", "%"+firstname+"%")
	}
	if lastname != "" {
		query = query.Where("last_name_th ILIKE ?", "%"+lastname+"%")
	}
	if middlename != "" {
		query = query.Where("middle_name_th ILIKE ?", "%"+middlename+"%")
	}
	if nationId != "" {
		query = query.Where("national_id ILIKE ?", "%"+nationId+"%")
	}
	if passportId != "" {
		query = query.Where("passport_id ILIKE ?", "%"+passportId+"%")
	}
	if dateOfBirth != "" {
		query = query.Where("birthday = ?", dateOfBirth)
	}
	if phone != "" {
		query = query.Where("phone ILIKE ?", "%"+phone+"%")
	}
	if email != "" {
		query = query.Where("email ILIKE ?", "%"+email+"%")
	}
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&patients).Error; err != nil {
		return nil, err
	}
	patientsEntity := make([]entity.Patient, len(patients))
	for i, patient := range patients {
		patientsEntity[i] = *patientToEntity(&patient)
	}
	return patientsEntity, nil
}

func (r *patientRepository) GeneratePatientCode() (string, error) {
	var lastPatient Patient
	if err := r.db.Order("created_at desc").First(&lastPatient).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "PT0000001", nil
		}
		return "", err
	}
	lastCode := lastPatient.PatientHn
	var lastNumber int
	_, err := fmt.Sscanf(lastCode, "PT%07d", &lastNumber)
	if err != nil {
		return "", err
	}
	newCode := fmt.Sprintf("PT%07d", lastNumber+1)
	return newCode, nil
}
