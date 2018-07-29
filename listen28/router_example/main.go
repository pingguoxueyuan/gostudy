package main

import "github.com/gin-gonic/gin"

func index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "index",
	})
}

func search(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "search",
	})
}

func support(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "support",
	})
}

func blog(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "blog",
	})
}

func about(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "about",
	})
}

func contact(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "contact",
	})
}

func main() {
	//Default返回一个默认的路由引擎
	router := gin.Default()

	router.POST("/", index)
	router.POST("/search", search)
	router.POST("/support", support)
	router.POST("/blog/:post", blog)
	router.POST("/about", about)
	router.POST("/contact", contact)

	router.Run(":8080")
}
