package middleware

import (
	"my_mange_system/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个JWTAuth的中间件
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 通过http header中的token解析来认证
		token := ctx.Request.Header.Get("token")
		if token == "" {
			common.Error(ctx, http.StatusUnauthorized, "token校验失败")
			return
		}
		// 初始化一个JWT对象实例，并根据结构体方法来解析token
		j := common.NewJWT()
		// 解析token中包含的相关信息(有效载荷)
		claims, err := j.ParserToken(token)
		if err != nil {
			common.Error(ctx, http.StatusUnauthorized, "token校验失败")
			return
		}
		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		ctx.Set("claims", claims)

	}
}
