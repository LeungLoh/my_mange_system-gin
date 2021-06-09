package api

import (
	"my_mange_system/server"
	"net/http"

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
		Offset:   1,
		Limit:    10,
	}
	ctx.ShouldBindQuery(&userinfo)
	users, total := server.GetUsetList(userinfo.Username, userinfo.Roleid, (userinfo.Offset-1)*userinfo.Limit, userinfo.Limit)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":  200,
		"data":  users,
		"total": total,
	})
}

func UserDelete(c *gin.Context) {
}
