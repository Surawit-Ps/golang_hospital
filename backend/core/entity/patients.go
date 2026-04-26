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
