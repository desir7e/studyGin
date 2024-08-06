package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	})
}
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	})
}
