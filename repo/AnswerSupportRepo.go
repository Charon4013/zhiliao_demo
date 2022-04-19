/**
 * @Author: Pan
 * @Date: 2022/4/10 16:50
 */

package repo

import (
	"fmt"
	"zhiliao_mvc_demo/datasource"
	"zhiliao_mvc_demo/model"
)

type AnswerSupportRepo interface {
	CreateSupport(uid int64, aid int64) bool
	GiveSupport(uid int64, aid int64) bool
	CancelSupport(uid int64, aid int64) bool
	CountSupport(aid int64) int64
	IsUserSupported(uid int64, aid int64) bool
	IsUserRecorded(uid int64, aid int64) bool
}

type answerSupportRepo struct{}

func NewAnswerSupportRepo() AnswerSupportRepo {
	return &answerSupportRepo{}
}

func (asr answerSupportRepo) CreateSupport(uid int64, aid int64) bool {
	var supportEngine = datasource.Init("AnswerSupport")
	defer supportEngine.Close()

	s := model.AnswerSupport{
		Uid:    uid,
		Aid:    aid,
		Status: false,
	}

	_, err := supportEngine.Table("answer_support").Insert(s)

	if err != nil {
		fmt.Println("Insert error: ", err)
		return false
	}
	return true
}

func (asr answerSupportRepo) GiveSupport(uid int64, aid int64) bool {
	var supportEngine = datasource.Init("AnswerSupport")
	defer supportEngine.Close()

	s := model.AnswerSupport{
		Status: true,
	}
	//fmt.Println("s: ", s)

	_, err := supportEngine.Table("answer_support").Where("uid = ? AND aid = ?", uid, aid).Cols("status").Update(&s)

	if err != nil {
		fmt.Println("Insert error: ", err)
		return false
	}
	return true
}

func (asr answerSupportRepo) CancelSupport(uid int64, aid int64) bool {
	var supportEngine = datasource.Init("AnswerSupport")
	defer supportEngine.Close()

	s := model.AnswerSupport{
		Status: false,
	}

	_, err := supportEngine.Table("answer_support").Where("uid = ? and aid = ?", uid, aid).Cols("status").Update(&s)

	if err != nil {
		return false
	}
	return true
}

func (asr answerSupportRepo) CountSupport(aid int64) int64 {
	support := new(model.AnswerSupport)
	var supportEngine = datasource.Init("AnswerSupport")
	defer supportEngine.Close()

	counts, err := supportEngine.Table("answer_support").Where("aid = ? AND status = ?", aid, true).Count(support)
	if err != nil {
		return 0
	}
	return counts
}

func (asr answerSupportRepo) CountSupportByUid(uid int64) int64 {
	support := new(model.AnswerSupport)
	var supportEngine = datasource.Init("AnswerSupport")
	defer supportEngine.Close()

	counts, err := supportEngine.Table("answer_support").Where("aid = ? AND status = ?", uid, true).Count(support)
	if err != nil {
		return 0
	}
	return counts
}

func (asr answerSupportRepo) IsUserSupported(uid int64, aid int64) bool {
	var supportEngine = datasource.Init("AnswerSupport")
	defer supportEngine.Close()

	exist, err := supportEngine.Table("answer_support").Where("uid = ? AND aid = ? AND status = ?", uid, aid, true).Exist()
	if err != nil {
		return false
	}
	return exist
}

func (asr answerSupportRepo) IsUserRecorded(uid int64, aid int64) bool {
	var supportEngine = datasource.Init("AnswerSupport")
	defer supportEngine.Close()

	exist, err := supportEngine.Table("answer_support").Where("uid = ? AND aid = ?", uid, aid).Exist()
	if err != nil {
		return false
	}
	return exist
}
