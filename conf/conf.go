package conf

import (
	"my_mange_system/model"
	"my_mange_system/util"
	"os"
)

func Init() {

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
}
