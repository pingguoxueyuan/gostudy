package dal

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pingguoxueyuan/gostudy/blogger/model"
)

var (
	DB *sqlx.DB
)

func Init(dns string) error {
	var err error
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

func InsertArticle(article *model.Article) (articleId int64, err error) {

	sqlstr := "insert into article(content, title, username)values(?,?,?)"
	result, err := DB.Exec(sqlstr, article.Content, article.Title, article.Username)
	if err != nil {
		return
	}

	articleId, err = result.LastInsertId()
	return
}

func GetArticleList(pageNum, pageSize int) (articleList []*model.Article, err error) {

	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}

	sqlstr := `select 
					id, content, title, view_count,
					 create_time, comment_count, username
				from 
					article 
				where 
					status = 1
				limit ?, ?`

	err = DB.Select(&articleList, sqlstr, pageNum, pageSize)
	return
}

func InsertComment(comment *model.Comment) (err error) {

	if comment == nil {
		err = fmt.Errorf("invalid parameter")
		return
	}

	tx, err := DB.Beginx()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	sqlstr := `insert 
					into comment(
						Content, Username, ArticleId					
					)
				values (
						?, ?, ?
				)`

	_, err = tx.Exec(sqlstr, comment.Content, comment.Username, comment.ArticleId)
	if err != nil {
		return
	}

	sqlstr = `  update 
					article 
				set 
					comment_count = comment_count + 1
				where
					id = ?`

	_, err = tx.Exec(sqlstr, comment.ArticleId)
	if err != nil {
		return
	}

	err = tx.Commit()
	return
}

func UpdateViewCount(articleId int64) (err error) {

	sqlstr := ` update 
					article 
				set 
					comment_count = comment_count + 1
				where
					id = ?`

	_, err = DB.Exec(sqlstr, articleId)
	if err != nil {
		return
	}

	return
}

func GetCommentList(articleId int64, pageNum, pageSize int) (commentList []*model.Comment, err error) {

	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}

	sqlstr := `select 
						id, content, username, create_time
					from 
						comment 
					where 
						article_id = ? and 
						status = 1
					limit ?, ?`

	err = DB.Select(&commentList, sqlstr, articleId, pageNum, pageSize)
	return
}
