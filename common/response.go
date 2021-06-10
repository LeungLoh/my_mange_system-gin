package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnAuthorized(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"status": http.StatusUnauthorized,
		"msg":    "token校验失败",
		"data":   nil,
	})
	ctx.Abort()
}

func BadRequest(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusBadRequest,
		"msg":    msg,
		"data":   nil,
	})
}

func InternalServerError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"status": http.StatusInternalServerError,
		"msg":    msg,
		"data":   nil,
	})
}

func Success(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    msg,
		"data":   data,
	})
}
