package api

import (
	"my_mange_system/server"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {

}

func UserLogin(c *gin.Context) {

}

func UserLogout(c *gin.Context) {

}

func UserList(ctx *gin.Context) {
	users := server.GetUsetList()
	// fmt.Println(users)
	Success(ctx, users, "test")
}

func UserDelete(c *gin.Context) {

}
