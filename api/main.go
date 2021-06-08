package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//请求成功的时候 使用该方法返回信息
func Success(ctx *gin.Context, data interface{}, msg interface{}) {
	fmt.Println(data)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}

//请求失败的时候, 使用该方法返回信息
func Failed(ctx *gin.Context, data interface{}, msg interface{}) {
	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"code": http.StatusBadRequest,
		"msg":  msg,
		"data": data,
	})
}
