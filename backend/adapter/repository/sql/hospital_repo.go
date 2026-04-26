package repository

import (
	"backend/core/entity"
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

func (r *hospitalRepository) CreateHospital(hospital *entity.Hospital) (*entity.Hospital, error) {
	hospital.Id = uuid.New().String()
	hospital.CreatedAt = time.Now()
	hospital.UpdatedAt = time.Now()
	if err := r.db.Create(hospital).Error; err != nil {
		return nil, err
	}
	return hospital, nil
}

func (r *hospitalRepository) GetHospitalById(id string) (*entity.Hospital, error) {
	var hospital entity.Hospital
	if err := r.db.First(&hospital, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &hospital, nil
}

func (r *hospitalRepository) GetHospitalByCode(code string) (*entity.Hospital, error) {
	var hospital entity.Hospital
	if err := r.db.First(&hospital, "hospital_code = ?", code).Error; err != nil {
		return nil, err
	}

	return &hospital, nil
}


