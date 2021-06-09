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

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func UserRegister(ctx *gin.Context) {

}

func PostUserLogin(ctx *gin.Context) {
	var user User

	if ctx.ShouldBind(&user) == nil {
		server.GenerateToken(ctx, user.Username)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "用户数据解析失败",
			"data":   nil,
		})
	}
}

func UserLogout(ctx *gin.Context) {

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

func UserDelete(ctx *gin.Context) {
}
