package main

import "github.com/gin-gonic/gin"

func testHandle(c *gin.Context) {
	c.Request.Cookie()
	c.JSON(200, gin.H{
		"message": "test",
	})
}

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", testHandle)

	//r.Run() // listen and serve on 0.0.0.0:8080
	r.Run(":9090")
}
