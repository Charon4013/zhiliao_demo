/**
 * @Author: Pan
 * @Date: 2022/4/13 15:57
 */

package repo

import (
	"fmt"
	"zhiliao_mvc_demo/datasource"
	"zhiliao_mvc_demo/model"
)

type CommentSupportRepo interface {
	CreateSupport(uid int64, cid int64) bool
	GiveSupport(uid int64, cid int64) bool
	CancelSupport(uid int64, cid int64) bool
	CountSupport(cid int64) int64
	IsUserSupported(uid int64, cid int64) bool
}

type commentSupportRepo struct{}

func NewCommentSupportRepo() CommentSupportRepo {
	return &commentSupportRepo{}
}

func (csr commentSupportRepo) CreateSupport(uid int64, cid int64) bool {
	var supportEngine = datasource.Init("CommentSupport")
	defer supportEngine.Close()

	s := model.CommentSupport{
		Uid:    uid,
		Cid:    cid,
		Status: false,
	}

	_, err := supportEngine.Table("comment_support").Insert(s)

	if err != nil {
		fmt.Println("Insert error: ", err)
		return false
	}
	return true
}

func (csr commentSupportRepo) GiveSupport(uid int64, cid int64) bool {
	var supportEngine = datasource.Init("CommentSupport")
	defer supportEngine.Close()

	s := model.CommentSupport{
		Status: true,
	}
	//fmt.Println("s: ", s)

	_, err := supportEngine.Table("comment_support").Where("uid = ? AND cid = ?", uid, cid).Cols("status").Update(&s)

	if err != nil {
		fmt.Println("Insert error: ", err)
		return false
	}
	return true
}

func (csr commentSupportRepo) CancelSupport(uid int64, cid int64) bool {
	var supportEngine = datasource.Init("CommentSupport")
	defer supportEngine.Close()

	s := model.CommentSupport{
		Status: false,
	}

	_, err := supportEngine.Table("comment_support").Where("uid = ? and cid = ?", uid, cid).Cols("status").Update(&s)

	if err != nil {
		return false
	}
	return true
}

func (csr commentSupportRepo) CountSupport(cid int64) int64 {
	support := new(model.AnswerSupport)
	var supportEngine = datasource.Init("CommentSupport")
	defer supportEngine.Close()

	counts, err := supportEngine.Table("comment_support").Where("cid = ? AND status = ?", cid, true).Count(support)
	if err != nil {
		return 0
	}
	return counts
}

func (csr commentSupportRepo) IsUserSupported(uid int64, cid int64) bool {
	var supportEngine = datasource.Init("CommentSupport")
	defer supportEngine.Close()

	exist, err := supportEngine.Table("comment_support").Where("uid = ? AND cid = ? AND status = ?", uid, cid, true).Exist()
	if err != nil {
		return false
	}
	return exist
}
