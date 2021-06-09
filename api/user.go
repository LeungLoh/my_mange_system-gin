package api

import (
	"my_mange_system/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserListParams struct {
	Username string `form:"username"`
	Roleid   int    `form:"roleid"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit"`
}

type UserLoginParams struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func UserRegister(ctx *gin.Context) {

}

func UserLogin(ctx *gin.Context) {
	var userloginparams UserLoginParams
	if ctx.ShouldBind(&userloginparams) == nil {
		if server.CheckOutUser(userloginparams.Username, userloginparams.Password) == true {
			server.GenerateToken(ctx, userloginparams.Username)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "账号或密码错误",
				"data":   nil,
			})
		}

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "用户数据解析失败",
			"data":   nil,
		})
	}
}

func UserInfo(ctx *gin.Context) {
	var userinfoparams UserLoginParams
	ctx.ShouldBindQuery(&userinfoparams)
	username, roleid := server.GetUserinfo(userinfoparams.Username)
	ctx.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "获取信息成功",
		"data":   gin.H{"username": username, "roleid": roleid},
	})
}
func UserList(ctx *gin.Context) {
	var userlistparams = UserListParams{
		Username: "",
		Roleid:   0,
		Offset:   1,
		Limit:    10,
	}
	ctx.ShouldBindQuery(&userlistparams)
	users, total := server.GetUsetList(userlistparams.Username, userlistparams.Roleid, (userlistparams.Offset-1)*userlistparams.Limit, userlistparams.Limit)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":  200,
		"data":  users,
		"total": total,
	})
}

func UserDelete(ctx *gin.Context) {
}

func UserLogout(ctx *gin.Context) {

}
