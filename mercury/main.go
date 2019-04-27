package main

import (
	"github.com/DeanThompson/ginpprof"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/logger"
	"github.com/pingguoxueyuan/gostudy/mercury/controller/account"
	"github.com/pingguoxueyuan/gostudy/mercury/controller/answer"
	"github.com/pingguoxueyuan/gostudy/mercury/controller/category"
	"github.com/pingguoxueyuan/gostudy/mercury/controller/comment"
	"github.com/pingguoxueyuan/gostudy/mercury/controller/favorite"
	"github.com/pingguoxueyuan/gostudy/mercury/controller/question"
	"github.com/pingguoxueyuan/gostudy/mercury/dal/db"
	"github.com/pingguoxueyuan/gostudy/mercury/filter"
	"github.com/pingguoxueyuan/gostudy/mercury/id_gen"
	maccount "github.com/pingguoxueyuan/gostudy/mercury/middleware/account"
	"github.com/pingguoxueyuan/gostudy/mercury/util"
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
	//err = maccount.InitSession("redis", "localhost:6379")
	err = maccount.InitSession("memory", "")
	return
}

func initFilter() (err error) {

	err = filter.Init("./data/filter.dat.txt")
	if err != nil {
		logger.Error("init filter failed, err:%v", err)
		return
	}

	logger.Debug("init filter succ")
	return
}

func main() {
	router := gin.Default()
	config := make(map[string]string)
	config["log_level"] = "debug"
	logger.InitLogger("console", config)

	err := initDb()
	if err != nil {
		panic(err)
	}

	err = initFilter()
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

	err = util.InitKafka("localhost:9092")
	if err != nil {
		panic(err)
	}
	err = util.InitKafkaConsumer("localhost:9092", "mercury_topic",
		func(message *sarama.ConsumerMessage) {
			logger.Debug("receive from kafka, msg:%#v", message)
		})
	if err != nil {
		panic(err)
		return
	}

	ginpprof.Wrapper(router)
	initTemplate(router)
	router.POST("/api/user/register", account.RegisterHandle)
	router.POST("/api/user/login", account.LoginHandle)
	router.GET("/api/category/list", category.GetCategoryListHandle)
	router.POST("/api/ask/submit", maccount.AuthMiddleware, question.QuestionSubmitHandle)
	router.GET("/api/question/list", category.GetQuestionListHandle)
	router.GET("/api/question/detail", question.QuestionDetailHandle)
	router.GET("/api/answer/list", answer.AnswerListHandle)
	router.POST("/api/answer/commit", maccount.AuthMiddleware, answer.AnswerCommitHandle)

	//评论模块
	commentGroup := router.Group("/api/comment/", maccount.AuthMiddleware)
	//commentGroup := router.Group("/api/comment/")
	commentGroup.POST("/post_comment", comment.PostCommentHandle)
	commentGroup.POST("/post_reply", comment.PostReplyHandle)
	commentGroup.GET("/list", comment.CommentListHandle)
	commentGroup.GET("/reply_list", comment.ReplyListHandle)
	commentGroup.POST("/like", comment.LikeHandle)

	//收藏模块路由
	favoriteGroup := router.Group("/api/favorite/")
	favoriteGroup.POST("/add_dir", favorite.AddDirHandle)
	favoriteGroup.POST("/add", favorite.AddFavoriteHandle)
	favoriteGroup.GET("/dir_list", favorite.DirListHandle)
	favoriteGroup.GET("/list", favorite.FavoriteListHandle)

	router.Run(":9090")
}
