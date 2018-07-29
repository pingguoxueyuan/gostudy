package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
