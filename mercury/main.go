package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/mercury/controller/account"
)

func main() {
	router := gin.Default()

	/*
		dns := "root:root@tcp(localhost:3306)/blogger?parseTime=true"
		err := db.Init(dns)
		if err != nil {
			panic(err)
		}*/

	ginpprof.Wrapper(router)
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")

	router.GET("/user/login", account.LoginViewHandle)
	router.GET("/user/register", account.RegisterViewHandle)
	router.Run(":8080")
}
