package repository

import (
	"backend/core/entity"
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"
)

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *staffRepository {
	return &staffRepository{db: db}
}

type Staff struct {
	Id         string    `gorm:"primaryKey"`
	HospitalId string    `gorm:"index"`
	Username   string    `gorm:"uniqueIndex"`
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func staffToEntity(staff *Staff) *entity.Staff {
	return &entity.Staff{
		Id:         staff.Id,
		HospitalId: staff.HospitalId,
		Username:   staff.Username,
		Password:   staff.Password,
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
		CreatedAt:  staff.CreatedAt,
		UpdatedAt:  staff.UpdatedAt,
	}
}

func (r *staffRepository) CreateStaff(staff *entity.Staff) (*entity.Staff, error) {
	staff.Id = uuid.New().String()
	staff.CreatedAt = time.Now()
	staff.UpdatedAt = time.Now()
	if err := r.db.Create(staff).Error; err != nil {
		return nil, err
	}
	return staff, nil
}

func (r *staffRepository) GetStaffById(id string) (*entity.Staff, error) {
	var staff entity.Staff
	if err := r.db.First(&staff, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}

func (r *staffRepository) GetStaffByUsername(username string) (*entity.Staff, error) {
	var staff entity.Staff
	if err := r.db.First(&staff, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}