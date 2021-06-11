package common

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Result struct {
	Httpcode int         `json:"httpcode"`
	Data     interface{} `json:"data"`
	Msg      string      `json:"msg"`
}

func Success(ctx *gin.Context, httpcode int, msg string, data interface{}, token string) {
	ctx.JSON(httpcode, gin.H{
		"status": httpcode,
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

func SetSession(ctx *gin.Context, key string, value interface{}) {
	session := sessions.Default(ctx)
	session.Set(key, value)
	err := session.Save()
	if err != nil {
		fmt.Println(err)
	}
}

func GetSession(ctx *gin.Context, key string) interface{} {
	session := sessions.Default(ctx)
	value := session.Get(key)
	return value
}
