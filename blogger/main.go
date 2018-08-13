package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/blogger/controller"
	"github.com/pingguoxueyuan/gostudy/blogger/dal/db"
)

func main() {
	router := gin.Default()

	dns := "root:root@tcp(localhost:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}

	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")

	router.GET("/", controller.IndexHandle)
	//发布文章页面
	router.GET("/article/new/", controller.NewArticle)
	//文章提交接口
	router.POST("/article/submit/", controller.ArticleSubmit)
	//文章详情页
	router.GET("/article/detail/", controller.ArticleDetail)

	//文件上传接口
	router.POST("/upload/file/", controller.UploadFile)

	//留言页面
	router.GET("/leave/new/", controller.LeaveNew)
	//关于我页面
	router.GET("/about/me/", controller.AboutMe)
	router.Run(":8080")
}
