package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	RoleId   string `gorm:"column:roleid"`
	Status   bool   `gorm:"column:status"`
}

func (u *User) TableName() string {
	return "user"
}
