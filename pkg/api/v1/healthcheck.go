package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary get list healthcheck
// @Description healthcheck
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 500 {string} string "Internal server error"
// @Router /healthcheck [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
