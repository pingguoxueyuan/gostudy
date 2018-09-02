package logic

import (
	"fmt"
	"math"

	"github.com/pingguoxueyuan/gostudy/blogger/dal/db"
	"github.com/pingguoxueyuan/gostudy/blogger/model"
)

func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {

LABEL:
	for _, article := range articleInfoList {
		categoryId := article.CategoryId
		for _, id := range ids {
			if id == categoryId {
				continue LABEL
			}
		}

		ids = append(ids, categoryId)
	}
	return
}

func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//1. 从数据库中，获取文章列表
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		fmt.Printf("1 get article list failed, err:%v\n", err)
		return
	}

	if len(articleInfoList) == 0 {
		return
	}

	categoryIds := getCategoryIds(articleInfoList)
	//2. 从数据库中，获取文章对应的分类列表
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("2 get category list failed, err:%v\n", err)
		return
	}

	//聚合数据
	for _, article := range articleInfoList {
		fmt.Printf("content:%s\n", article.Summary)
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}

	return
}


func GetArticleRecordListById(categoryId, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//1. 从数据库中，获取文章列表
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		fmt.Printf("1 get article list failed, err:%v\n", err)
		return
	}

	if len(articleInfoList) == 0 {
		return
	}

	categoryIds := getCategoryIds(articleInfoList)
	//2. 从数据库中，获取文章对应的分类列表
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("2 get category list failed, err:%v\n", err)
		return
	}

	//聚合数据
	for _, article := range articleInfoList {
		fmt.Printf("content:%s\n", article.Summary)
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}

	return
}


func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {

	//1. 获取文章的信息
	articleDetail, err = db.GetArticleDetail(articleId)
	if err != nil {
		return
	}

	//2. 获取文章对应的分类信息
	category, err := db.GetCategoryById(articleDetail.ArticleInfo.CategoryId)
	if err != nil {
		return
	}

	articleDetail.Category = *category
	return
}

func GetRelativeAricleList(articleId int64) (articleList []*model.RelativeArticle, err error) {

	articleList, err = db.GetRelativeArticle(articleId)
	return
}

func GetPrevAndNextArticleInfo(articleId int64) (prevArticle *model.RelativeArticle,
	nextArticle *model.RelativeArticle, err error) {

	prevArticle, err = db.GetPrevArticleById(articleId)
	if err != nil {
		//打印一个警告日志
	}

	nextArticle, err = db.GetNextArticleById(articleId)
	if err != nil {
		//打印一个警告日志
	}

	return
}

/*
	content := c.PostForm("content")
	author := c.PostForm("author")
	categoryId := c.PostForm("category_id")
	title := c.PostForm("title")
*/
func InsertArticle(content, author, title string, categoryId int64) (err error) {

	//1. 构造一个ArticleDetail的类
	articleDetail := &model.ArticleDetail{}
	articleDetail.Content = content
	articleDetail.Username = author
	articleDetail.Title = title
	articleDetail.ArticleInfo.CategoryId = categoryId

	contentUtf8 := []rune(content)
	minLength := int(math.Min(float64(len(contentUtf8)), 128.0))
	articleDetail.Summary = string([]rune(content)[:minLength])

	id, err := db.InsertArticle(articleDetail)
	fmt.Printf("insert article succ, id:%d, err:%v\n", id, err)
	return
}
