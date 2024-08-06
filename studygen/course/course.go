package course

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// validator 逻辑验证
type course struct {
	Name     string `form:"name"  json:"name" binding:"required,alphaunicode"`
	Teacher  string `form:"teacher"   json:"teacher" binding:"required,alphaunicode"`
	Duration int    `form:"duration"  json:"duration" binding:"required,numeric"`
}

func Add(c *gin.Context) {
	//定义结构体指针
	req := &course{}
	//参数绑定
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, req)
}
func Get(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func Update(c *gin.Context) {

	req := &course{}
	//参数绑定
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, req)
}

func Delete(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
