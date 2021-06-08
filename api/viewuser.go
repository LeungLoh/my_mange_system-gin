package api

import (
	"my_mange_system/server"

	"github.com/gin-gonic/gin"
)

type Userinfo struct {
	Username string `form:"username"`
	Roleid   int    `form:"roleid"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit"`
}

func UserRegister(c *gin.Context) {

}

func UserLogin(c *gin.Context) {

}

func UserLogout(c *gin.Context) {

}

func UserList(ctx *gin.Context) {
	var userinfo = Userinfo{
		Username: "",
		Roleid:   0,
		Offset:   0,
		Limit:    10,
	}
	ctx.ShouldBindQuery(&userinfo)
	users := server.GetUsetList(userinfo.Username, userinfo.Roleid, userinfo.Offset, userinfo.Limit)

	Success(ctx, users, "test")
}

func UserDelete(c *gin.Context) {
}
