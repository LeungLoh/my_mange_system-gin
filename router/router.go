package router

import (
	"encoding/gob"
	"my_mange_system/api"
	"my_mange_system/middleware"
	"my_mange_system/model"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gob.Register(model.User{})
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.Sessions("secret"))
	r.Use(middleware.Response())
	v1 := r.Group("/api/v1")
	{
		v1.GET("user/info", api.UserInfo)
		v1.PUT("user/info", api.UserChangePassword)
		login := v1.Group("user/login")
		login.POST("", api.UserLogin)
		login.DELETE("", api.UserLogout)
		userlist := v1.Group("user/list")
		// userlist.Use(middleware.JWTAuth())
		userlist.GET("", api.UserList)
		userlist.DELETE("", api.UserDelete)
		userlist.PUT("", api.UserUpdate)

	}

	return r
}
