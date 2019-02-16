package db

import (
	"github.com/pingguoxueyuan/gostudy/logger"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
)

func CreateFavoriteDir(favoriteDir *common.FavoriteDir) (err error) {
	tx, err := DB.Beginx()
	if err != nil {
		logger.Error("create favorite dir failed, favorite dir:%#v, err:%v", favoriteDir, err)
		return
	}

	//先查询相同的dir_name是否存在
	var dirCount int64
	sqlstr := `select count(dir_id) from favorite_dir where user_id=? and dir_name=?`
	err = tx.Get(&dirCount, sqlstr, favoriteDir.UserId, favoriteDir.DirName)
	if err != nil {
		logger.Error("select dir_name failed, err:%v, favoriteDir:%#v", err, favoriteDir)
		return
	}

	if dirCount > 0 {
		tx.Rollback()
		err = ErrRecordExists
		return
	}

	sqlstr = `	insert 
				into favorite_dir (
						user_id, dir_id, dir_name
					)
				values (
						?, ?, ?
				)`

	_, err = tx.Exec(sqlstr, favoriteDir.UserId, favoriteDir.DirId, favoriteDir.DirName)
	if err != nil {
		logger.Error("insert favorite_dir failed, favorite_dir:%#v err:%v", favoriteDir, err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("insert favorite_dir failed, favorite_dir:%#v err:%v", favoriteDir, err)
		tx.Rollback()
		return
	}
	return
}

func CreateFavorite(favorite *common.Favorite) (err error) {
	tx, err := DB.Beginx()
	if err != nil {
		logger.Error("create favorite dir failed, favorite :%#v, err:%v", favorite, err)
		return
	}

	//先查询相同的dir_name是否存在
	var favoriteCount int64
	sqlstr := `select count(answer_id) from favorite where user_id=? and dir_id=?`
	err = tx.Get(&favoriteCount, sqlstr, favorite.UserId, favorite.DirId)
	if err != nil {
		logger.Error("select dir_name failed, err:%v, favorite:%#v", err, favorite)
		return
	}

	if favoriteCount > 0 {
		tx.Rollback()
		err = ErrRecordExists
		return
	}

	sqlstr = `	insert 
				into favorite (
						user_id, dir_id, answer_id
					)
				values (
						?, ?, ?
				)`

	_, err = tx.Exec(sqlstr, favorite.UserId, favorite.DirId, favorite.AnswerId)
	if err != nil {
		logger.Error("insert favorite failed, favorite:%#v err:%v", favorite, err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("insert favorite failed, favorite:%#v err:%v", favorite, err)
		tx.Rollback()
		return
	}
	return
}

func GetFavoriteDirList(userId int64) (favoriteDirList []*common.FavoriteDir, err error) {

	sqlstr := `select dir_id, dir_name, count
	 			from favorite_dir
	  			where user_id=?`
	err = DB.Select(&favoriteDirList, sqlstr, userId)
	if err != nil {
		logger.Error("select favorite dir failed, err:%v", err)
		return
	}

	return
}

func GetFavoriteList(userId, dirId, offset, limit int64) (favoriteList []*common.Favorite, err error) {

	sqlstr := `select dir_id, user_id, answer_id
					 from favorite
					  where user_id=? and dir_id=? limit ?, ?`
	err = DB.Select(&favoriteList, sqlstr, userId, dirId, offset, limit)
	if err != nil {
		logger.Error("select favorite list failed, err:%v", err)
		return
	}

	return
}
