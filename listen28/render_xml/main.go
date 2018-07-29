package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/moreXML", func(c *gin.Context) {
		// You also can use a struct
		type MessageRecord struct {
			Name    string
			Message string
			Number  int
		}

		var msg MessageRecord
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.XML(http.StatusOK, msg)
	})
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
