package routers

import (
	"github.com/gin-gonic/gin"
	"studygen/user"
)

func initUser(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		v1.GET("/user", user.Get)
		v1.POST("/user", user.Add)
	}
	v2 := group.Group("/v2")
	{
		v2.GET("/user", user.GetV2)
		v2.POST("/user", user.AddV2)
	}

}
