package routers

import (
	"github.com/gin-gonic/gin"
)

// HealthGET is delegated to inform the caller about the status of the service
func HealthGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}
