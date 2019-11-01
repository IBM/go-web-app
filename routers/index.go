package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index will expose the homepage
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// NotFoundError will manage the 404 error
func NotFoundError(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", nil)
}

// InternalServerError will manage the 505 error
func InternalServerError(c *gin.Context) {
	c.HTML(http.StatusOK, "500.html", nil)
}
