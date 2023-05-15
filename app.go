package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// 默认监听 8080 端口
	err := r.Run()

	if err != nil {
		return
	}
}
