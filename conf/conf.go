package conf

import (
	"fmt"
	"my_mange_system-gin/model"
	"my_mange_system-gin/util"
)

func Init() {

	// 设置日志级别
	util.BuildLogger("info")

	// 连接数据库
	username := "root"          //账号
	password := "123456"        //密码
	host := "127.0.0.1"         //数据库地址，可以是Ip或者域名
	port := 3306                //数据库端口
	Dbname := "my_mange_system" //数据库名
	timeout := "10s"            //连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	model.Database(dsn)
}
