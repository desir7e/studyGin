package routers

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	api := r.Group("/api")
	//课程接口
	initCourse(api)
	//用户接口
	initUser(api)
	//登录
	initLogin(api)
}
