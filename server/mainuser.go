package server

import (
	"my_mange_system/model"
)

type UserList struct {
	Username string `json:"username"`
	Roleid   string `json:"roleid"`
	Userid   uint   `json:"userid"`
}

func GetUsetList() []UserList {
	var users []model.User
	var new_users []UserList
	model.DB.Find(&users).Offset(0).Limit(10)
	for _, user := range users {
		row := UserList{
			Username: user.Username,
			Roleid:   user.RoleId,
			Userid:   user.Model.ID,
		}

		new_users = append(new_users, row)
	}
	return new_users
}
