package routers

import (
	"github.com/gin-gonic/gin"
	"studygen/course"
)

func initCourse(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{ //路由传参
		v1.GET("/course/:id", course.Get)
		v1.POST("/course", course.Add)
		v1.PUT("/course", course.Update)
		v1.DELETE("/course", course.Delete)

	}

}
