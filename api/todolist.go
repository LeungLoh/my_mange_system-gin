package api

import (
	"my_mange_system-gin/common"
	"my_mange_system-gin/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TodolistHandleParams struct {
	ID     uint   `form:"id"`
	Offset int    `form:"offset"`
	Limit  int    `form:"limit"`
	Title  string `form:"title"`
	Status bool   `form:"status"`
}

func GetTodoList(ctx *gin.Context) {
	var todolist []model.Todolist
	var total int64
	var params = UserHandleParams{
		Offset: 1,
		Limit:  5,
	}

	ctx.ShouldBindQuery(&params)
	DB := model.DB.Model(&model.Todolist{})
	user := GetUserInfo(ctx)
	DB.Where("userid = ?", user.ID)
	DB.Count(&total)
	DB.Order("createtime desc").Limit(params.Limit).Offset(params.Offset - 1).Find(&todolist)
	res := common.Result{Httpcode: http.StatusOK, Msg: "获取信息成功", Data: gin.H{"todolist": todolist, "total": total}}
	ctx.Set("Res", res)
	ctx.Next()
}

func InsertTodoList(ctx *gin.Context) {
	var params TodolistHandleParams
	var res common.Result
	DB := model.DB.Model(&model.Todolist{})
	if ctx.ShouldBind(&params) == nil {
		user := GetUserInfo(ctx)
		todo := model.Todolist{Createtime: time.Now(), Title: params.Title, Status: false, UserId: user.ID}
		DB.Create(&todo)
		res = common.Result{Httpcode: http.StatusOK, Msg: "创建成功"}
	} else {
		res = common.Result{Httpcode: http.StatusBadRequest, Err: "用户数据解析失败"}
	}
	ctx.Set("Res", res)
	ctx.Next()
}

func UpdateTodoList(ctx *gin.Context) {
	var params TodolistHandleParams
	var res common.Result
	DB := model.DB.Model(&model.Todolist{})
	if ctx.ShouldBind(&params) == nil {
		user := GetUserInfo(ctx)
		DB.Where(map[string]interface{}{"id": params.ID, "userid": user.ID}).Update("status", params.Status)
		res = common.Result{Httpcode: http.StatusOK, Msg: "修改成功"}
	} else {
		res = common.Result{Httpcode: http.StatusBadRequest, Err: "用户数据解析失败"}
	}
	ctx.Set("Res", res)
	ctx.Next()
}

func DeleteTodoList(ctx *gin.Context) {
	var params TodolistHandleParams
	var res common.Result
	DB := model.DB.Model(&model.Todolist{})
	if ctx.ShouldBind(&params) == nil {
		user := GetUserInfo(ctx)
		DB.Where(map[string]interface{}{"id": params.ID, "userid": user.ID}).Delete(&model.Todolist{})
		res = common.Result{Httpcode: http.StatusOK, Msg: "删除成功"}
	} else {
		res = common.Result{Httpcode: http.StatusBadRequest, Err: "用户数据解析失败"}
	}
	ctx.Set("Res", res)
	ctx.Next()
}

func GetUserInfo(ctx *gin.Context) model.User {
	user := common.GetSession(ctx, "user")
	if user == nil {
		res := common.Result{Httpcode: http.StatusInternalServerError, Err: "无法获取用户信息"}
		ctx.Set("Res", res)
		ctx.Next()
		return model.User{}
	}
	userinfo := user.(model.User)
	return userinfo
}
