package common

import (
	_ "github.com/gin-gonic/gin"
	"net/http"
)

func Success(data any, msg ...string) map[string]any {
	message := resolveMsg(msg)
	return map[string]any{
		"code": http.StatusOK,
		"msg":  message,
		"data": data,
	}
}

func Error(msg ...string) map[string]any {
	message := resolveErrorMsg(msg)
	return map[string]any{
		"code": http.StatusInternalServerError,
		"msg":  message,
		"data": nil,
	}
}

func SuccessArray(data ...any) map[string]any {
	return map[string]any{
		"code": http.StatusOK,
		"msg":  "操作成功",
		"data": data,
	}
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
