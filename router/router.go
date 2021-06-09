package router

import (
	"my_mange_system/api"
	"my_mange_system/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	// r.Use(middleware.JWTAuth())
	v1 := r.Group("/api/v1")
	{
		v1.POST("user/login", api.PostUserLogin)

		auth := v1.Group("")
		auth.Use(middleware.JWTAuth())
		{
			v1.GET("user/list", middleware.JWTAuth(), api.UserList)
		}

	}
	return r
}
