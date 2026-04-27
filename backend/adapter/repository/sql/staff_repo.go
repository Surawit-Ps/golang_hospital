package repository

import (
	"backend/core/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *staffRepository {
	return &staffRepository{db: db}
}

type Staff struct {
	Id         string `gorm:"primaryKey"`
	HospitalId string `gorm:"index;column:hospital_id"`
	Username   string `gorm:"uniqueIndex:idx_username_hospital;index:idx_username_hospital"`
	Password   string
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func staffToEntity(staff *Staff) *entity.Staff {
	return &entity.Staff{
		Id:         staff.Id,
		HospitalId: staff.HospitalId,
		Username:   staff.Username,
		Password:   staff.Password,
		Role:       staff.Role,
		CreatedAt:  staff.CreatedAt,
		UpdatedAt:  staff.UpdatedAt,
	}
}

func entityToStaff(staff *entity.Staff) *Staff {
	return &Staff{
		Id:         staff.Id,
		HospitalId: staff.HospitalId,
		Username:   staff.Username,
		Password:   staff.Password,
		Role:       staff.Role,
		CreatedAt:  staff.CreatedAt,
		UpdatedAt:  staff.UpdatedAt,
	}
}

func (r *staffRepository) CreateStaff(staff *entity.Staff) error {
	staff.Id = uuid.New().String()
	staff.CreatedAt = time.Now()
	staff.UpdatedAt = time.Now()
	staff.Role = "staff"
	staffModel := entityToStaff(staff)
	if err := r.db.Create(staffModel).Error; err != nil {
		return err
	}
	return nil
}

func (r *staffRepository) GetStaffById(id string) (*entity.Staff, error) {
	var staff Staff
	if err := r.db.First(&staff, "id = ?", id).Error; err != nil {
		return nil, err
	}
	staffEntity := staffToEntity(&staff)
	return staffEntity, nil
}

func (r *staffRepository) GetStaffByUsernameAndHospitalId(username, hospitalId string) (*entity.Staff, error) {
	var staff Staff
	query := r.db.Where("username = ?", username)

	// If hospitalId is provided, filter by it
	if hospitalId != "" {
		// If hospitalId looks like a HospitalNumber (e.g., "HOSP0001"), resolve to UUID first
		if len(hospitalId) == 8 && hospitalId[:4] == "HOSP" {
			var hospital Hospital
			if err := r.db.Where("hospital_number = ?", hospitalId).First(&hospital).Error; err != nil {
				return nil, err
			}
			query = query.Where("hospital_id = ?", hospital.Id)
		} else {
			query = query.Where("hospital_id = ?", hospitalId)
		}
	}

	if err := query.First(&staff).Error; err != nil {
		return nil, err
	}
	staffEntity := staffToEntity(&staff)
	return staffEntity, nil
}

func (r *staffRepository) GetUser() (*entity.Staff, error) {
	var staff Staff
	if err := r.db.First(&staff).Error; err != nil {
		return nil, err
	}
	staffEntity := staffToEntity(&staff)
	return staffEntity, nil
}
