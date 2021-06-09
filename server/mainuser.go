package server

import (
	"fmt"
	"my_mange_system/model"
)

type UserList struct {
	Username string `json:"username"`
	Roleid   int    `json:"roleid"`
	Userid   uint   `json:"userid"`
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
	fmt.Println(total)
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
