package dal

import (
	"fmt"
	"time"

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

type Article struct {
	Id           int64     `db:"id"`
	Content      string    `db:"content"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CreateTime   time.Time `db:"create_time"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
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

func GetArticleList(pageNum, pageSize int) (articleList []*Article, err error) {

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
