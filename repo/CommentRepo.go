/**
 * @Author: Pan
 * @Date: 2022/4/12 12:58
 */

package repo

import (
	"fmt"
	"zhiliao_mvc_demo/datasource"
	"zhiliao_mvc_demo/model"
)

var comment model.Comment

type CommentRepo interface {
	CreateComment(c model.Comment) (int, model.Comment)
	DeleteComment(cid int64) bool
	//GetCommentByUid(cid int64) (int, model.Comment)
	GetAllCommentByAid(aid int64) (int, []model.Comment)
	CountCommentByAid(aid int64) int64
	// No Modify
}

type commentRepo struct{}

func NewCommentRepo() CommentRepo {
	return &commentRepo{}
}

func (cr commentRepo) CreateComment(c model.Comment) (int, model.Comment) {
	var commentEngine = datasource.Init("Comment")
	defer commentEngine.Close()

	_, err := commentEngine.Table("comment").Insert(c)
	if err != nil {
		fmt.Println("Insert error: ", err)
		return -1, model.Comment{}
	}

	return 0, c
}

func (cr commentRepo) DeleteComment(cid int64) bool {
	var commentEngine = datasource.Init("Comment")
	defer commentEngine.Close()

	_, err := commentEngine.Table("comment").ID(cid).Delete(&comment)
	if err != nil {
		return false
	}
	return true
}

//func (cr commentRepo) GetCommentByUid(uid int64) (int, model.Comment) {
//	defer commentEngine.Close()
//
//	affected ,err := commentEngine.Table("comment").Where("uid = ?", uid).Get(&comment)
//	if err != nil || affected == false{
//		return -1, model.Comment{}
//	}
//	return 0, comment
//}

func (cr commentRepo) GetAllCommentByAid(aid int64) (int, []model.Comment) {
	var commentList []model.Comment

	var commentEngine = datasource.Init("Comment")
	defer commentEngine.Close()

	err := commentEngine.Table("comment").Where("aid = ?", aid).OrderBy("created").Find(&commentList)
	if err != nil || len(commentList) == 0 {
		fmt.Println("sql err or commentList = 0: ", err)
		return -1, nil
	}
	return 0, commentList
}

func (cr commentRepo) CountCommentByAid(aid int64) int64 {
	comment := new(model.Comment)
	var commentEngine = datasource.Init("Comment")
	defer commentEngine.Close()

	counts, err := commentEngine.Table("comment").Where("aid = ?", aid).Count(comment)
	if err != nil {
		return 0
	}
	return counts
}
