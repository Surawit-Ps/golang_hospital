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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	createdHospital, err := h.hospitalService.RegisterHospital(&hospital)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, createdHospital)
}

func (h *hospitalHandler) GetHospitalInfo(c *gin.Context) {
	id := c.Param("id")
	hospital, err := h.hospitalService.GetHospitalInfo(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "hospital not found"})
		return
	}
	c.JSON(200, hospital)
}

