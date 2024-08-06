package main

import (
	"studygen/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.InitRouters(r)
	//注册路由分组  用于分模块

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
