package main

import "github.com/gin-gonic/gin"

func login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}

func read(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}

func submit(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}

func main() {
	//Default返回一个默认的路由引擎
	router := gin.Default()

	// Simple group: v1
	//   /v1/login
	//   /v1/submit
	//   /v1/read
	v1 := router.Group("/v1")
	{
		v1.POST("/login", login)
		v1.POST("/submit", submit)
		v1.POST("/read", read)
	}

	// Simple group: v2
	//   /v2/login
	//   /v2/submit
	//   /v2/read
	v2 := router.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
		v2.POST("/read", read)
	}

	router.Run(":8080")
}
