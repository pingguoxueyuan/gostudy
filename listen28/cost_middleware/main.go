package main

import (
	"log"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		//可以设置一些公共参数
		c.Set("example", "12345")
		//等其他中间件先执行
		c.Next()
		//获取耗时
		latency := time.Since(t)
		log.Printf("total cost time:%d us", latency/1000)
	}
}

func main() {
	//r := gin.New()
	r := gin.Default()
	r.Use(StatCost())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
