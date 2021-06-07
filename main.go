package main

import (
	"github.com/gin-gonic/gin"
	"my_mange_system/router"
)

func main() {
	r := router.NewRouter()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8000") // 监听并在 0.0.0.0:8080 上启动服务
}
