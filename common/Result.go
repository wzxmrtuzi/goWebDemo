package common

import "net/http"

func Success(data interface{}, msg ...string) map[string]interface{} {
	message := resolveMsg(msg)
	return map[string]interface{}{
		"code": http.StatusOK,
		"msg":  message,
		"data": data,
	}
}
func SuccessArray(data ...interface{}) map[string]interface{} {
	return map[string]interface{}{
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
