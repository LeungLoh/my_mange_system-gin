package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Httpcode int         `json:"httpcode"`
	Data     interface{} `json:"data"`
	Msg      string      `json:"msg"`
}

func Success(ctx *gin.Context, msg string, data interface{}, token string) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    msg,
		"data":   data,
		"token":  token,
	})
}
func Error(ctx *gin.Context, httpcode int, msg string) {
	ctx.JSON(httpcode, gin.H{
		"status": httpcode,
		"msg":    msg,
	})
}
