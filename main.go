package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init() {
	fmt.Println("初始化")
	strconv.Atoi("123")
}

func _string(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello")
}

func _json(c *gin.Context) {
	var user struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Password string `json:"-"`
	}
	user.Name = "张三"
	user.Age = 12
	user.Password = "123456"
	// res := make(map[string]interface{})
	// res["msg"] = "操作成功"
	// res["data"] = user
	res := map[string]interface{}{
		"msg":  "操作成功",
		"data": user,
	}
	c.JSON(http.StatusOK, res)
}

func _xml(c *gin.Context) {
	res := gin.H{"msg": "成功", "data": gin.H{"id": 1, "name": "糊涂"}}
	c.XML(http.StatusOK, res)
}
func _yaml(c *gin.Context) {
	res := gin.H{"msg": "成功", "data": gin.H{"id": 1, "name": "糊涂"}}
	c.YAML(http.StatusOK, res)
}

func _html(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("web/**")
	router.GET("/string", _string)
	router.GET("/ping", _json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.Run(":80")
}
