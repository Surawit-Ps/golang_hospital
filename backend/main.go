package main

import (
	"backend/core/entity"
	"backend/core/service"
	"backend/adapter/handler"
	"backend/adapter/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize GORM with SQLite
	db, err := gorm.Open(sqlite.Open("hospital.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate the database schema
	db.AutoMigrate(&entity.Hospital{}, &entity.Staff{}, &entity.Patient{})
	// Initialize repositories
	hospitalRepo := repository.NewHospitalRepository(db)
	staffRepo := repository.NewStaffRepository(db)
	patientRepo := repository.NewPatientRepository(db)
	// Initialize services
	hospitalService := service.NewHospitalService(hospitalRepo)
	staffService := service.NewStaffService(staffRepo, hospitalRepo)
	patientService := service.NewPatientService(patientRepo, hospitalRepo)
	// Initialize handlers
	hospitalHandler := handler.NewHospitalHandler(hospitalService)
	staffHandler := handler.NewStaffHandler(staffService)
	patientHandler := handler.NewPatientHandler(patientService)
	
}