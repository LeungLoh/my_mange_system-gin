package model

type User struct {
	ID            uint   `gorm:"primarykey"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	RoleId        int    `gorm:"column:roleid"`
	City          string `gorm:"column:city"`
	LastLoginTime int64  `gorm:"column:lastlogintime"`
	Status        bool   `gorm:"column:status"`
}

func (u *User) TableName() string {
	return "user"
}
