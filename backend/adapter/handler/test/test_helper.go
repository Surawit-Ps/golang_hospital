package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Helper types for testing
type Param struct {
	Key   string
	Value string
}

type ParamError struct {
	Error string
}

// createTestContext creates a test Gin context
func createTestContext(w http.ResponseWriter, req *http.Request) *gin.Context {
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	return c
}
