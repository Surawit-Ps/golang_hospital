package repository

import (
	"backend/core/entity"
	"fmt"
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type hospitalRepository struct {
	db *gorm.DB
}

func NewHospitalRepository(db *gorm.DB) *hospitalRepository {
	return &hospitalRepository{db: db}
}

type Hospital struct {
	Id           string    `gorm:"primaryKey"`
	Name         string
	HospitalNumber string    `gorm:"uniqueIndex"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func hospitalToEntity(hospital *Hospital) *entity.Hospital {
	return &entity.Hospital{
		Id:           hospital.Id,
		Name:         hospital.Name,
		HospitalNumber: hospital.HospitalNumber,
		CreatedAt:    hospital.CreatedAt,
		UpdatedAt:    hospital.UpdatedAt,
	}
}

func entityToHospital(hospital *entity.Hospital) *Hospital {
	return &Hospital{
		Id:           hospital.Id,
		Name:         hospital.Name,
		HospitalNumber: hospital.HospitalNumber,
		CreatedAt:    hospital.CreatedAt,
		UpdatedAt:    hospital.UpdatedAt,
	}
}

func (r *hospitalRepository) CreateHospital(hospital *entity.Hospital) (*entity.Hospital, error) {
	hospital.Id = uuid.New().String()
	hospital.CreatedAt = time.Now()
	hospital.UpdatedAt = time.Now()
	code, err := r.GenerateHospitalCode()
	if err != nil {
		return nil, err
	}
	hospital.HospitalNumber = code
	formHospital := entityToHospital(hospital)
	if err := r.db.Create(formHospital).Error; err != nil {
		return nil, err
	}
	return hospital, nil
}

func (r *hospitalRepository) GetHospitalById(id string) (*entity.Hospital, error) {
	var hospital Hospital
	if err := r.db.First(&hospital, "id = ? OR hospital_number = ?", id, id).Error; err != nil {
		return nil, err
	}
	return hospitalToEntity(&hospital), nil
}

func (r *hospitalRepository) GetAllHospitals() ([]entity.Hospital, error) {
	var hospitals []Hospital
	if err := r.db.Find(&hospitals).Error; err != nil {
		return nil, err
	}
	hospitalsEntity := make([]entity.Hospital, len(hospitals))
	for i, hospital := range hospitals {
		hospitalsEntity[i] = *hospitalToEntity(&hospital)
	}
	return hospitalsEntity, nil
}

func (r *hospitalRepository) GenerateHospitalCode() (string, error) {
	var lastHospital Hospital
	if err := r.db.Order("created_at desc").First(&lastHospital).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "HOSP0001", nil
		}
		return "", err
	}
	lastCode := lastHospital.HospitalNumber
	var lastNumber int
	_, err := fmt.Sscanf(lastCode, "HOSP%04d", &lastNumber)
	if err != nil {
		return "", err
	}
	newCode := fmt.Sprintf("HOSP%04d", lastNumber+1)
	return newCode, nil
}

