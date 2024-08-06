package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	})
}
func GetV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	})
}
