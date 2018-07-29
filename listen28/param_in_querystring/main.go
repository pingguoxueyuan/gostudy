package main

import "github.com/gin-gonic/gin"

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "少林")
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message":  "pong",
			"username": username,
			"address":  address,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
