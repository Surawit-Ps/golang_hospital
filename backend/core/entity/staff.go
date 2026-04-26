package entity

import "time"

type Staff struct {
	Id string 
	Username string
	Password string
	HospitalId string
	CreatedAt time.Time
	UpdatedAt time.Time
}