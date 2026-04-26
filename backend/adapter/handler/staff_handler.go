package handler

import (
	"backend/core/entity"
	"backend/core/ports"
	"github.com/gin-gonic/gin"
)

type StaffHandler struct {
	staffService ports.StaffService
}

func NewStaffHandler(staffService ports.StaffService) *StaffHandler {
	return &StaffHandler{
		staffService: staffService,
	}
}

func (h *StaffHandler) RegisterStaff(c *gin.Context) {
	var staff entity.Staff
	if err := c.ShouldBindJSON(&staff); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	createdStaff, err := h.staffService.RegisterStaff(&staff)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, createdStaff)
}

func (h *StaffHandler) LoginStaff(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	staff, err := h.staffService.LoginStaff(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}
	c.JSON(200, staff)
}

