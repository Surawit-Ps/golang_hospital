package handler

import (
	"backend/core/entity"
	"backend/core/ports"
	"github.com/gin-gonic/gin"
)

type hospitalHandler struct {
	hospitalService ports.HospitalService
	
}

func NewHospitalHandler(hospitalService ports.HospitalService) *hospitalHandler {
	return &hospitalHandler{
		hospitalService: hospitalService,
	}
}

func (h *hospitalHandler) RegisterHospital(c *gin.Context) {
	var hospital entity.Hospital
	if err := c.ShouldBindJSON(&hospital); err != nil {
		handleError(c,err)
		return
	}
	createdHospital, err := h.hospitalService.RegisterHospital(&hospital)
	if err != nil {
		handleError(c,err)
		return
	}
	ResponseCreateSuccess(c, "Hospital registered successfully", createdHospital)
}

func (h *hospitalHandler) GetHospitalInfo(c *gin.Context) {
	id := c.Param("id")
	hospital, err := h.hospitalService.GetHospitalInfo(id)
	if err != nil {
		handleError(c, err)
		return
	}
	ResponseSuccess(c, hospital)
}

func (h *hospitalHandler) GetAllHospitals(c *gin.Context) {
	hospitals, err := h.hospitalService.GetAllHospitals()
	if err != nil {
		handleError(c, err)
		return
	}
	ResponseSuccess(c, hospitals)
}

