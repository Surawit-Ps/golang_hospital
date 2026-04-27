package entity

import "time"

type Staff struct {
	Id string 
	Username string
	Password string
	HospitalId string
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time
}