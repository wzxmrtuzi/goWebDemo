package controllers

import (
	"goWebDemo/entity"
	"goWebDemo/response"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)
	user.Id = "1"
	response.Success(c, user)
}