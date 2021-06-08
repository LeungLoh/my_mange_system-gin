package router

import (
	"my_mange_system/api"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		v1.POST("user/logout", api.UserLogout)
		v1.GET("user/list", api.UserList)
		v1.DELETE("user/delete", api.UserDelete)
	}
	return r
}
