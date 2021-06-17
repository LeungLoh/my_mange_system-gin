package api

import (
	"my_mange_system-gin/common"
	"my_mange_system-gin/model"
	"my_mange_system-gin/server"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserLoginParams struct {
	Username string `form:"username"`
	Password string `form:"password"`
	City     string `form:"city"`
}

type UserListParams struct {
	Username string `form:"username"`
	Roleid   int    `form:"roleid"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit"`
}

type UserHandleParams struct {
	Username string `form:"username"`
	Password string `form:"password"`
	UserId   string `form:"userid"`
	Roleid   string `form:"roleid"`
	City     string `form:"city"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit"`
}

type ChangePassword struct {
	Username    string `form:"username"`
	OldPassword string `form:"oldpassword"`
	NewPassword string `form:"newpassword"`
	UserId      int    `form:"userid"`
}

func UserLogin(ctx *gin.Context) {
	var params UserHandleParams
	var res common.Result
	if ctx.ShouldBind(&params) == nil {
		result, user := server.CheckOutUser(ctx, params.Username, params.Password)
		if result == true {
			server.UpdateLoginInfo(params.City, params.Username)
			res = common.Result{Httpcode: http.StatusOK, Msg: "登录成功", Data: gin.H{"username": user.Username, "userid": user.ID, "roleid": user.RoleId}}
		} else {
			res = common.Result{Httpcode: http.StatusNoContent, Err: "账号密码错误"}
		}
	} else {
		res = common.Result{Httpcode: http.StatusBadRequest, Err: "用户数据解析失败"}
	}
	ctx.Set("Res", res)
	ctx.Next()
}

func UserLogout(ctx *gin.Context) {
	s := sessions.Default(ctx)
	s.Clear()
	s.Save()
	res := common.Result{Httpcode: http.StatusOK, Msg: "退出登录成功"}
	ctx.Set("Res", res)
	ctx.Next()
}

func UserInfo(ctx *gin.Context) {
	user := common.GetSession(ctx, "user")
	if user == nil {
		res := common.Result{Httpcode: http.StatusInternalServerError, Err: "无法获取用户信息"}
		ctx.Set("Res", res)
		ctx.Next()
		return
	}
	userinfo := user.(model.User)
	data := gin.H{
		"username":      userinfo.Username,
		"roleid":        userinfo.RoleId,
		"city":          userinfo.City,
		"lastlogintime": time.Unix(userinfo.LastLoginTime, 0).Format("2006-01-02 15:04:05"),
	}

	res := common.Result{Httpcode: http.StatusOK, Msg: "获取信息成功", Data: data}
	ctx.Set("Res", res)
	ctx.Next()
}

func UserChangePassword(ctx *gin.Context) {
	var params ChangePassword
	var res common.Result
	user := common.GetSession(ctx, "user")
	if user == nil {
		res := common.Result{Httpcode: http.StatusInternalServerError, Err: "无法获取用户信息"}
		ctx.Set("Res", res)
		ctx.Next()
		return
	}
	userinfo := user.(model.User)
	if ctx.ShouldBind(&params) == nil {
		if uint(params.UserId) != userinfo.ID {
			res = common.Result{Httpcode: http.StatusBadRequest, Err: "只能修改自己用户信息"}
		} else {
			result, msg := server.ChangeUserPassword(uint(params.UserId), params.OldPassword, params.NewPassword)
			if result == true {
				res = common.Result{Httpcode: http.StatusOK, Msg: msg}
			} else {
				res = common.Result{Httpcode: http.StatusBadRequest, Err: msg}
			}
		}
	} else {
		res = common.Result{Httpcode: http.StatusBadRequest, Err: "用户数据解析失败"}
	}
	ctx.Set("Res", res)
	ctx.Next()
}

func UserList(ctx *gin.Context) {
	var params = UserHandleParams{
		Username: "",
		Roleid:   "0",
		Offset:   1,
		Limit:    10,
	}
	var res common.Result
	ctx.ShouldBindQuery(&params)
	roleid, err := strconv.Atoi(params.Roleid)
	if err != nil {
		res = common.Result{Httpcode: http.StatusBadRequest, Err: "用户数据解析失败"}
	} else {
		users, total := server.GetUsetList(params.Username, roleid, (params.Offset-1)*params.Limit, params.Limit)
		res = common.Result{Httpcode: http.StatusOK, Msg: "获取信息成功", Data: gin.H{"users": users, "total": total}}
	}
	ctx.Set("Res", res)
	ctx.Next()
}

func UserDelete(ctx *gin.Context) {
	var params UserHandleParams
	var res common.Result
	user := common.GetSession(ctx, "user")
	if user == nil {
		res := common.Result{Httpcode: http.StatusInternalServerError, Err: "无法获取用户信息"}
		ctx.Set("Res", res)
		ctx.Next()
		return
	}
	userinfo := user.(model.User)
	if userinfo.RoleId != 1 {
		res := common.Result{Httpcode: http.StatusUnauthorized, Err: "非管理员无法删除"}
		ctx.Set("Res", res)
		ctx.Next()
		return
	}
	if ctx.ShouldBind(&params) == nil {
		userids := strings.Split(params.UserId, ",")
		roleids := strings.Split(params.Roleid, ",")
		result, msg := server.DeleteUserList(userids, roleids, userinfo)
		if result == false {
			res = common.Result{Httpcode: http.StatusNoContent, Err: msg}
		} else {
			res = common.Result{Httpcode: http.StatusOK, Msg: msg}
		}
	} else {
		res = common.Result{Httpcode: http.StatusBadRequest, Err: "用户数据解析失败"}
	}

	ctx.Set("Res", res)
	ctx.Next()
}

func SystemInfo(ctx *gin.Context) {
	meminfo := server.ReadSystemInfo("meminfo")
	data := gin.H{
		"MemTotal":     meminfo["MemTotal"],
		"MemFree":      meminfo["MemFree"],
		"Buffers":      meminfo["Buffers"],
		"Cached":       meminfo["Cached"],
		"MemAvailable": meminfo["MemAvailable"],
	}
	res := common.Result{Httpcode: http.StatusOK, Msg: "获取信息成功", Data: data}
	ctx.Set("Res", res)
	ctx.Next()
}

func Registered(ctx *gin.Context) {
	var params UserHandleParams
	var res common.Result
	var total int64
	DB := model.DB.Model(&model.User{})
	if ctx.ShouldBind(&params) == nil {
		user := model.User{
			Username: params.Username,
			Password: params.Password,
		}
		DB.Where("username = ?", params.Username).Count(&total)
		if total > 0 {
			res = common.Result{Httpcode: http.StatusBadRequest, Err: "用户名已存在"}
		} else {
			DB.Create(&user)
			res = common.Result{Httpcode: http.StatusOK, Msg: "注册成功"}
		}

	} else {
		res = common.Result{Httpcode: http.StatusBadRequest, Err: "用户数据解析失败"}
	}
	ctx.Set("Res", res)
	ctx.Next()
}
