package middleware

import (
	"my_mange_system-gin/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		token, err := common.GenerateToken("leungloh")
		if err != nil {
			common.Error(ctx, http.StatusInternalServerError, "token创建失败")
		}
		temp, ok := ctx.Get("Res")
		if !ok {
			common.Error(ctx, http.StatusInternalServerError, "获取信息失败")
		}
		res := temp.(common.Result)
		if res.Err == "" {
			common.Success(ctx, res.Httpcode, res.Msg, res.Data, token)
		} else {
			common.Error(ctx, res.Httpcode, res.Err)
		}

	}
}
