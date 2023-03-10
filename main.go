package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"goWebDemo/response"
)

func InitServer(router *gin.Engine, address ...string) *http.Server {
	gin.SetMode(gin.ReleaseMode)

	addr := resolveAddress(address)
	return &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		return ":80"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}

func init() {
	res, err := strconv.Atoi("123")
	if err != nil {
		log.Println("错误", err)
	}
	log.Printf("初始化,res: %v", res)
}

func _string(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello")
}

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"-"`
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
	// res := map[string]interface{}{
	// 	"msg":  "操作成功",
	// 	"data": user,
	// }
	res := make([]User, 0)
	user1 := User{}
	user1.Name = "李四"
	user1.Age = 12
	user1.Password = "123456"
	user2 := User{}
	user2.Name = "张三"
	user2.Age = 23
	user2.Password = "123456ffff"
	res = append(res, user1, user2)
	response.Success(c, res)
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

func _redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}
func _query(c *gin.Context) {
	// /query?user=abc&age=1,2
	param1 := c.Query("user")
	param2 := c.QueryArray("age")
	param3, ok3 := c.GetQuery("age")
	param4 := c.Query("password")
	param5, ok5 := c.GetQuery("password")
	fmt.Println(param1)
	fmt.Println(param2)
	fmt.Println(param3, ok3)
	fmt.Println(param4)
	fmt.Println(param5, ok5)
}

func _param(c *gin.Context) {
	// /param/张三/123
	user := c.Param("user")
	age := c.Param("age")
	password := c.DefaultPostForm("password", "123456")
	response.SuccessArray(c, user, age, password)
}

func _form(c *gin.Context) {
	// form-data
	// name := c.PostForm("name")
	// nameArr := c.PostFormArray("name")
	// c.JSON(http.StatusOK, response.SuccessArray(name, nameArr, files, err, fileValue))

	// 接收所有参数和文件
	params, _ := c.MultipartForm()
	// 参数
	value := params.Value
	// 文件
	files := params.File
	response.SuccessArray(c, files, value)
}

func _raw(c *gin.Context) {
	// application/x-www-form-urlencoded
	raw, err := c.GetRawData()
	if err != nil {
		fmt.Println("失败")
	}
	contentType := c.GetHeader("content-type")
	response.SuccessArray(c, raw, contentType)
}

func _body(c *gin.Context) {
	var user struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	err := c.ShouldBind(&user)
	if err == nil {
		response.Success(c, user)
	} else {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
	}
}

func _header(c *gin.Context) {
	c.Header("token", "yyyyyyyyy")
	response.Success(c, nil)
}

var router = gin.Default()

func main() {

	router.LoadHTMLGlob("web/**")
	router.StaticFile("/dogImg", "static/dog.jpg")
	router.StaticFS("/fs", http.Dir("static/text"))
	router.StaticFileFS("/fileFs", "static/hello.text", http.Dir("static/text"))

	router.GET("/string", _string)
	router.GET("/json", _json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.GET("/baidu", _redirect)
	router.GET("/query", _query)
	router.GET("/param/:user/:age", _param)
	router.POST("/form", _form)
	router.POST("/raw", _raw)

	router.POST("/body", _body)

	router.POST("/header", _header)

	// 分组
	v1 := router.Group("api/v1")
	v1.GET("/json", _json)

	server := InitServer(router, ":9091")
	server.ListenAndServe()
}
