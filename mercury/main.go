package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/mercury/controller/account"
	"github.com/pingguoxueyuan/gostudy/mercury/dal/db"
	"github.com/pingguoxueyuan/gostudy/mercury/id_gen"
	maccount "github.com/pingguoxueyuan/gostudy/mercury/middleware/account"
)

func initTemplate(router *gin.Engine) {
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.Static("/css/", "./static/css/")
	router.Static("/fonts/", "./static/fonts/")
	router.Static("/img/", "./static/img/")
	router.Static("/js/", "./static/js/")
}

func initDb() (err error) {
	dns := "root:root@tcp(localhost:3306)/mercury?parseTime=true"
	err = db.Init(dns)
	if err != nil {
		return
	}

	return
}

func initSession() (err error) {
	err = maccount.InitSession("memory", "")
	return
}

func main() {
	router := gin.Default()

	err := initDb()
	if err != nil {
		panic(err)
	}

	err = initSession()
	if err != nil {
		panic(err)
	}

	err = id_gen.Init(1)
	if err != nil {
		panic(err)
	}

	ginpprof.Wrapper(router)
	initTemplate(router)
	router.POST("/api/user/register", account.RegisterHandle)
	router.POST("/api/user/login", account.LoginHandle)
	router.Run(":9090")
}
