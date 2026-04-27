package handler

import (
	"backend/core/entity"
	"backend/core/ports"
	"backend/pkg/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	patientService ports.PatientService
}

func NewPatientHandler(patientService ports.PatientService) *patientHandler {
	return &patientHandler{
		patientService: patientService,
	}
}

func (h *patientHandler) RegisterPatient(c *gin.Context) {
	user := c.MustGet("userID").(string)
	if user == "" {
		handleError(c, errs.ErrUnauthorized)
		return
	}
	var patient entity.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		handleError(c, err)
		return
	}
	err := h.patientService.RegisterPatient(&patient)
	if err != nil {
		handleError(c, err)
		return
	}
	ResponseCreateSuccess(c, "Patient registered successfully", patient)

}

func (h *patientHandler) GetPatientInfo(c *gin.Context) {
	user := c.MustGet("userID").(string)
	if user == "" {
		handleError(c, errs.ErrUnauthorized)
		return
	}
	id := c.Param("id")
	patient, err := h.patientService.GetPatientInfo(id)
	if err != nil {
		handleError(c, errs.ErrNotFound)
		return
	}
	ResponseSuccess(c, patient)
}

func (h *patientHandler) GetPatients(c *gin.Context) {
	user := c.MustGet("userID").(string)
	if user == "" {
		handleError(c, errs.ErrUnauthorized)
		return
	}
	hospitalId := c.MustGet("HospitalID").(string)
	firstname := c.Query("firstname")
	lastname := c.Query("lastname")
	middlename := c.Query("middlename")
	nationId := c.Query("nationId")
	passportId := c.Query("passportId")
	dateOfBirth := c.Query("dateOfBirth")
	phone := c.Query("phone")
	email := c.Query("email")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		handleError(c, errs.ErrInvalidPageNumber)
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		handleError(c, errs.ErrInvalidLimitNumber)
		return
	}
	patients, err := h.patientService.GetPatients(hospitalId, firstname, lastname, middlename, nationId, passportId, dateOfBirth, phone, email, pageInt, limitInt)
	if err != nil {
		handleError(c, err)
		return
	}
	ResponseSuccess(c, patients)
}
