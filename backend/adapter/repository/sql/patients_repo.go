package repository

import (
	"backend/core/entity"
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"
)

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *patientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) CreatePatient(patient *entity.Patient) (*entity.Patient, error) {
	patient.Id = uuid.New().String()
	patient.CreatedAt = time.Now()
	patient.UpdatedAt = time.Now()
	if err := r.db.Create(patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}

func (r *patientRepository) GetPatientById(id string) (*entity.Patient, error) {
	var patient entity.Patient
	if err := r.db.First(&patient, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) GetPatientsByHospitalId(hospitalId string) ([]*entity.Patient, error) {
	var patients []*entity.Patient
	if err := r.db.Where("hospital_id = ?", hospitalId).Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}