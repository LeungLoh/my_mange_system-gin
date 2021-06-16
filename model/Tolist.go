package model

import "time"

type Todolist struct {
	ID         uint      `gorm:"primary_key"`
	Createtime time.Time `gorm:"column:createtime"`
	Title      string    `gorm:"column:title"`
	Status     bool      `gorm:"column:status"`
	UserId     uint      `gorm:"column:userid"`
}

func (u *Todolist) TableName() string {
	return "todolist"
}
