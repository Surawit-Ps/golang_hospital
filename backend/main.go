package main

import (
	"backend/adapter/handler"
	repository "backend/adapter/repository/sql"
	"backend/core/entity"
	"backend/core/middleware"
	"backend/core/ports"
	"backend/core/service"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Get database configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "hospital"
	}

	// Create PostgreSQL connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Initialize GORM with PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Auto-migrate the database schema
	db.AutoMigrate(&repository.Hospital{}, &repository.Staff{}, &repository.Patient{})
	// Initialize repositories
	hospitalRepo := repository.NewHospitalRepository(db)
	staffRepo := repository.NewStaffRepository(db)
	patientRepo := repository.NewPatientRepository(db)
	// Initialize services
	hospitalService := service.NewHospitalService(hospitalRepo)
	staffService := service.NewStaffService(staffRepo)
	patientService := service.NewPatientService(patientRepo)
	// Initialize handlers
	hospitalHandler := handler.NewHospitalHandler(hospitalService)
	staffHandler := handler.NewStaffHandler(staffService)
	patientHandler := handler.NewPatientHandler(patientService)

	// Seed initial data
	allStaff, _ := staffRepo.GetUser()
	if allStaff == nil {
		setupData(staffService, patientService, hospitalService)
	}

	// Set up Gin router
	router := gin.Default()
	middleware.CORS()

	// Public routes - no auth required
	router.POST("/staff/login", staffHandler.LoginStaff)
	router.POST("/staff", staffHandler.RegisterStaff)

	// Protected routes - apply authorization middleware
	protected := router.Group("")
	protected.Use(middleware.Authorizes())
	{
		// Hospital routes
		protected.POST("/hospitals", hospitalHandler.RegisterHospital)
		protected.GET("/hospitals/:id", hospitalHandler.GetHospitalInfo)
		protected.GET("/hospitals", hospitalHandler.GetAllHospitals)
		// Staff routes
		protected.GET("/staff", staffHandler.GetUserLogin)
		// Patient routes
		protected.POST("/patients", patientHandler.RegisterPatient)
		protected.GET("/patients/:id", patientHandler.GetPatientInfo)
		protected.GET("/patients", patientHandler.GetPatients)
	}

	// Start the server
	router.Run(":8080")

}

func setupData(staffService ports.StaffService, patientService ports.PatientService, hospitalService ports.HospitalService) {
	hospital := []entity.Hospital{
		{Name: "Hospital A"},
		{Name: "Hospital B"},
		{Name: "Hospital C"},
		{Name: "Hospital D"},
		{Name: "Hospital E"},
	}

	// Store hospital IDs for staff assignment
	hospIds := []string{"HOSP0001", "HOSP0002", "HOSP0003", "HOSP0004", "HOSP0005"}
	for i, h := range hospital {
		h.HospitalNumber = hospIds[i]
		_, err := hospitalService.RegisterHospital(&h)
		if err != nil {
			panic("failed to seed hospital data")
		}
	}

	pass, err := middleware.HashPassword("123456")
	if err != nil {
		panic("failed to hash password")
	}

	// Use actual hospital IDs returned from registration
	staff := []entity.Staff{
		{Username: "staff1", Password: pass, HospitalId: hospIds[0]},
		{Username: "staff2", Password: pass, HospitalId: hospIds[1]},
		{Username: "staff3", Password: pass, HospitalId: hospIds[2]},
	}

	for _, s := range staff {
		err := staffService.RegisterStaff(&s)
		if err != nil {
			panic("failed to seed staff data: " + err.Error())
		}
	}

	patient := []entity.Patient{
		{HospitalId: hospIds[0], FirstNameTh: "นายเอ", LastNameTh: "สมชาย", MiddleNameTh: "ชาย", FirstNameEn: "A", LastNameEn: "Somchai", MiddleNameEn: "Chai", NationalId: "1111111111111", PassportId: "AA111111", Birthday: time.Date(1985, 3, 15, 0, 0, 0, 0, time.UTC), Phone: "0812345601", Email: "patientA@hospital.com", Gender: "M"},
		{HospitalId: hospIds[0], FirstNameTh: "นางบี", LastNameTh: "สมหญิง", MiddleNameTh: "หญิง", FirstNameEn: "B", LastNameEn: "Somying", MiddleNameEn: "Ying", NationalId: "2222222222222", PassportId: "BB222222", Birthday: time.Date(1990, 6, 20, 0, 0, 0, 0, time.UTC), Phone: "0812345602", Email: "patientB@hospital.com", Gender: "F"},
		{HospitalId: hospIds[0], FirstNameTh: "นายซี", LastNameTh: "วงศ์จันทร์", MiddleNameTh: "จันทร์", FirstNameEn: "C", LastNameEn: "Wongchan", MiddleNameEn: "Chan", NationalId: "3333333333333", PassportId: "CC333333", Birthday: time.Date(1988, 9, 10, 0, 0, 0, 0, time.UTC), Phone: "0812345603", Email: "patientC@hospital.com", Gender: "M"},
		{HospitalId: hospIds[0], FirstNameTh: "นางดี", LastNameTh: "สุขสวัสดิ์", MiddleNameTh: "สวัสดิ์", FirstNameEn: "D", LastNameEn: "Suksawat", MiddleNameEn: "Sat", NationalId: "4444444444444", PassportId: "DD444444", Birthday: time.Date(1992, 12, 5, 0, 0, 0, 0, time.UTC), Phone: "0812345604", Email: "patientD@hospital.com", Gender: "F"},
		{HospitalId: hospIds[1], FirstNameTh: "นายบีม", LastNameTh: "ศรีสวัสดิ์", MiddleNameTh: "สวัสดิ์", FirstNameEn: "Beam", LastNameEn: "Srisawat", MiddleNameEn: "Sat", NationalId: "5555555555555", PassportId: "BB555555", Birthday: time.Date(1987, 4, 12, 0, 0, 0, 0, time.UTC), Phone: "0812345605", Email: "patientB2@hospital.com", Gender: "M"},
		{HospitalId: hospIds[1], FirstNameTh: "นางบีม", LastNameTh: "ศรีสวัสดิ์", MiddleNameTh: "สวัสดิ์", FirstNameEn: "Beam", LastNameEn: "Srisawat", MiddleNameEn: "Sat", NationalId: "6666666666666", PassportId: "BB666666", Birthday: time.Date(1991, 7, 22, 0, 0, 0, 0, time.UTC), Phone: "0812345606", Email: "patientB3@hospital.com", Gender: "F"},
		{HospitalId: hospIds[1], FirstNameTh: "นายดำ", LastNameTh: "ดำเนิน", MiddleNameTh: "เนิน", FirstNameEn: "Dam", LastNameEn: "Damnerng", MiddleNameEn: "Nern", NationalId: "7777777777777", PassportId: "DD777777", Birthday: time.Date(1989, 11, 30, 0, 0, 0, 0, time.UTC), Phone: "0812345607", Email: "patientD2@hospital.com", Gender: "M"},
		{HospitalId: hospIds[2], FirstNameTh: "นายสมพงษ์", LastNameTh: "ศรีสกุล", MiddleNameTh: "สกุล", FirstNameEn: "Sompong", LastNameEn: "Srisukul", MiddleNameEn: "Sukul", NationalId: "8888888888888", PassportId: "AA888888", Birthday: time.Date(1986, 2, 14, 0, 0, 0, 0, time.UTC), Phone: "0812345608", Email: "patientA3@hospital.com", Gender: "M"},
		{HospitalId: hospIds[2], FirstNameTh: "นางสมหญิง", LastNameTh: "ศรีสวัสดิ์", MiddleNameTh: "สวัสดิ์", FirstNameEn: "Somying", LastNameEn: "Srisawat", MiddleNameEn: "Sat", NationalId: "9999999999999", PassportId: "AA999999", Birthday: time.Date(1993, 8, 8, 0, 0, 0, 0, time.UTC), Phone: "0812345609", Email: "patientA4@hospital.com", Gender: "F"},
		{HospitalId: hospIds[2], FirstNameTh: "นายศรี", LastNameTh: "ศรีสุข", MiddleNameTh: "สุข", FirstNameEn: "Sri", LastNameEn: "Srisuk", MiddleNameEn: "Suk", NationalId: "1010101010101", PassportId: "BB101010", Birthday: time.Date(1984, 5, 18, 0, 0, 0, 0, time.UTC), Phone: "0812345609", Email: "patientB3@hospital.com", Gender: "M"},
	}
	for _, p := range patient {
		err := patientService.RegisterPatient(&p)
		if err != nil {
			panic("failed to seed patient data: " + err.Error())
		}
	}
}
