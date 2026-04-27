package entity

import "time"

type Patient struct {
	Id string
	HospitalId string
	PatientHn string
	NationalId string
	PassportId string
	FirstNameTh string
	MiddleNameTh string
	LastNameTh string
	FirstNameEn string
	MiddleNameEn string
	LastNameEn string
	Birthday time.Time
	Phone string
	Email string
	Gender string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PatientResponse struct {
	FirstNameTh string `json:"first_name_th"`
	MiddleNameTh string `json:"middle_name_th"`
	LastNameTh string `json:"last_name_th"`
	FirstNameEn string `json:"first_name_en"`
	MiddleNameEn string `json:"middle_name_en"`
	LastNameEn string `json:"last_name_en"`
	Birthday time.Time `json:"date_of_birth"`
	PatientHn string `json:"patient_hn"`
	NationalId string `json:"national_id"`
	PassportId string `json:"passport_id"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Gender string `json:"gender"`
}