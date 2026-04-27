package errs

import "errors"

var (
	ErrHospitalNotFound = errors.New("hospital not found")
	ErrHospitalAlreadyExists = errors.New("hospital already exists")
	ErrInvalidHospitalData = errors.New("invalid hospital data")
	ErrDatabaseError = errors.New("database error")
	ErrInternalServerError = errors.New("internal server error")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound = errors.New("data not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidPageNumber = errors.New("invalid page number")
	ErrInvalidLimitNumber = errors.New("invalid limit number")
)