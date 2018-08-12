package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pingguoxueyuan/gostudy/blogger/model"
)

func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {

	if article == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}

	sqlstr := `insert into 
					article(content, summary, title, username, 
						category_id, view_count, comment_count)
				values(?, ?, ?, ?, ?, ?, ?)`

	result, err := DB.Exec(sqlstr, article.Content, article.Summary,
		article.Title, article.Username, article.Category.CategoryId,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}

	articleId, err = result.LastInsertId()
	return
}

func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {

	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}

	sqlstr := `select 
						id, summary, title, view_count,
						 create_time, comment_count, username, category_id
					from 
						article 
					where 
						status = 1
					order by create_time desc
					limit ?, ?`

	err = DB.Select(&articleList, sqlstr, pageNum, pageSize)
	return
}
