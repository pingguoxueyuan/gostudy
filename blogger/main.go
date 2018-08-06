package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	err := loadUploadConfig()
	if err != nil {
		panic("load upload config failed")
	}

	router.Static("/static", "./static")
	router.LoadHTMLGlob("views/*")

	router.GET("/", indexHandle)
	router.GET("/post", postHandle)
	router.GET("/upload/config", uploadConfigHandle)

	router.Run(":8080")
}
