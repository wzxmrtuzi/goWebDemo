package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(c *gin.Context, data any, msg ...string) {
	message := resolveMsg(msg)
	c.JSON(http.StatusOK, Result{http.StatusOK, message, data})
}

func Error(c *gin.Context, msg ...string) {
	message := resolveErrorMsg(msg)
	c.JSON(http.StatusOK, Result{http.StatusInternalServerError, message, nil})
}

func SuccessArray(c *gin.Context, data ...any) {
	c.JSON(http.StatusOK, Result{http.StatusOK, "操作成功", data})
}

func resolveMsg(msg []string) string {
	switch len(msg) {
	case 0:
		return "操作成功"
	case 1:
		return msg[0]
	default:
		panic("too many parameters")
	}
}
func resolveErrorMsg(msg []string) string {
	switch len(msg) {
	case 0:
		return "操作失败"
	case 1:
		return msg[0]
	default:
		panic("too many parameters")
	}
}
