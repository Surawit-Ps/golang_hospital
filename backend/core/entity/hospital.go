package entity

import "time"

type Hospital struct {
	Id string
	Name string
	HospitalCode string
	CreatedAt time.Time
	UpdatedAt time.Time
}