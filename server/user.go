package server

import (
	"my_mange_system/middleware"
	"my_mange_system/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserList struct {
	Username string `json:"username"`
	Roleid   int    `json:"roleid"`
	Userid   uint   `json:"userid"`
}

func CheckOutUser(username string, password string) bool {
	var user model.User
	DB := model.DB.Model(&model.User{})
	DB.Where("username = ?", username).First(&user)
	if user.Password == password {
		return true
	}
	return false
}
func GetUsetList(username string, roleid int, offset int, limit int) ([]UserList, int64) {
	var users []model.User
	var new_users []UserList
	var total int64
	DB := model.DB.Model(&model.User{})
	if username != "" {
		DB = DB.Where("username LIKE ?", "%"+username+"%")
	}
	if roleid > 0 {
		DB = DB.Where("roleid = ?", roleid)
	}
	DB.Count(&total)
	DB.Limit(limit).Offset(offset).Find(&users)
	for _, user := range users {
		row := UserList{
			Username: user.Username,
			Roleid:   user.RoleId,
			Userid:   user.ID,
		}

		new_users = append(new_users, row)
	}
	return new_users, total
}

func GenerateToken(ctx *gin.Context, username string) {
	jwt := middleware.NewJWT()
	claims := middleware.NewCustomClaims(username)
	token, err := jwt.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登陆成功",
		"data":   gin.H{"token": token},
	})
	return
}

func GetUserinfo(username string) (string, int) {
	var user model.User
	DB := model.DB.Model(&model.User{})
	DB.Where("username = ?", username).First(&user)
	return user.Username, user.RoleId
}
