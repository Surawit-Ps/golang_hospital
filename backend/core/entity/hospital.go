package entity

import "time"

type Hospital struct {
	Id string
	Name string
	HospitalNumber string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HospitalResponse struct {
	Name string `json:"name"`
	HospitalNumber string `json:"hospital_number"`
}