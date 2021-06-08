package main

import (
	"my_mange_system/conf"
	"my_mange_system/router"
)

func main() {
	conf.Init()
	app := router.NewRouter()
	app.Run(":8000") // 监听并在 0.0.0.0:8080 上启动服务
}
