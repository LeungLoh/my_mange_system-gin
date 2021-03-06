package router

import (
	"encoding/gob"
	"my_mange_system-gin/api"
	"my_mange_system-gin/middleware"
	"my_mange_system-gin/model"

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

		user := v1.Group("user")
		user.POST("registered", api.Registered)
		user.POST("login", api.UserLogin)
		user.DELETE("logout", middleware.JWTAuth(), api.UserLogout)

		userinfo := v1.Group("user/info")
		userinfo.Use(middleware.JWTAuth())
		userinfo.GET("", api.UserInfo)
		userinfo.PUT("", api.UserChangePassword)

		userlist := v1.Group("user/list")
		userlist.Use(middleware.JWTAuth())
		userlist.GET("", api.UserList)
		userlist.DELETE("", api.UserDelete)

		systeminfo := v1.Group("system")
		systeminfo.Use(middleware.JWTAuth())
		systeminfo.GET("info", api.SystemInfo)

		todolist := v1.Group("todolist")
		todolist.Use(middleware.JWTAuth())
		todolist.GET("", api.GetTodoList)
		todolist.POST("", api.InsertTodoList)
		todolist.PUT("", api.UpdateTodoList)
		todolist.DELETE("", api.DeleteTodoList)

	}

	return r
}
