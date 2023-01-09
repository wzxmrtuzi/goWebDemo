package api

import (
	"goWebDemo/controllers"

	"github.com/gin-gonic/gin"
)

func UserApi(router *gin.RouterGroup) {
	v1 := router.Group("/")
	{
		v1.POST("/book", controllers.AddUser)
	}
}
