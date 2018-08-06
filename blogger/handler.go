package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var (
	uploadConfig map[string]interface{}
)

func indexHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"title": "Posts",
	})
}

func postHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "views/post.html", gin.H{
		"title": "Posts",
	})
}

func uploadConfigHandle(c *gin.Context) {

	c.JSON(http.StatusOK, uploadConfig)
}

func loadUploadConfig() (err error) {
	filename := "./config/upload.json"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	err = json.Unmarshal(data, &uploadConfig)
	fmt.Printf("unmarshal failed, err:%v\n", err)
	return
}
