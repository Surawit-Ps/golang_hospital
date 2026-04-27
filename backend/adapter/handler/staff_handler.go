package handler

import (
	"backend/core/entity"
	"backend/core/middleware"
	"backend/core/ports"
	"backend/pkg/errs"
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
	user := c.MustGet("userID").(string)
	if user == "" {
		handleError(c, errs.ErrUnauthorized)
		return
	}
	var staff entity.Staff
	if err := c.ShouldBindJSON(&staff); err != nil {
		handleError(c, err)
		return
	}
	err := h.staffService.RegisterStaff(&staff)
	if err != nil {
		handleError(c, err)
		return
	}
	ResponseCreateSuccess(c, "Staff registered successfully", staff)
}

func (h *StaffHandler) LoginStaff(c *gin.Context) {
	var loginData struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		HospitalId string `json:"hospital_id"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		handleError(c, err)
		return
	}
	token, data, err := h.staffService.LoginStaff(loginData.Username, loginData.Password, loginData.HospitalId)
	if err != nil {
		handleError(c, errs.ErrInvalidCredentials)
		return
	}
	c.Set("userID", data.Id)
	c.Set("HospitalID", data.HospitalId)
	c.Set("role", data.Role)
	middleware.SetCookies(c, token)
	ResponseSuccess(c, token)
}

func (h *StaffHandler) GetUserLogin(c *gin.Context) {
	id := c.MustGet("userID").(string)
	if id == "" {
		handleError(c, errs.ErrUnauthorized)
		return
	}

	user, err := h.staffService.GetUser(id)
	if err != nil {
		handleError(c, errs.ErrNotFound)
		return
	}
	ResponseSuccess(c, user)
}
