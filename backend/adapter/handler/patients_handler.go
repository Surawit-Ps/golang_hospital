package handler

import (
	"backend/core/entity"
	"backend/core/ports"
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
	var patient entity.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	createdPatient, err := h.patientService.RegisterPatient(&patient)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, createdPatient)
}

func (h *patientHandler) GetPatientInfo(c *gin.Context) {
	id := c.Param("id")
	patient, err := h.patientService.GetPatientInfo(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "patient not found"})
		return
	}
	c.JSON(200, patient)
}

func (h *patientHandler) GetPatientsByHospital(c *gin.Context) {
	hospitalId := c.Param("hospitalId")
	patients, err := h.patientService.GetPatientsByHospital(hospitalId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, patients)
}

