/**
 * @Author: Pan
 * @Date: 2022/4/13 16:01
 */

package service

import (
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/repo"
)

type CommentSupportService interface {
	Create(uid int64, cid int64) (result model.Result)
	Add(uid int64, cid int64) (result model.Result)
	Cancel(uid int64, cid int64) (result model.Result)
	CountSupport(cid int64) int64
	HasUserSupported(uid int64, cid int64) bool
}

type commentSupportService struct{}

func NewCommentSupportService() CommentSupportService {
	return &commentSupportService{}
}

var commentSupportRepo = repo.NewCommentSupportRepo()

func (ass commentSupportService) Create(uid int64, cid int64) (result model.Result) {
	err := commentSupportRepo.CreateSupport(uid, cid)

	if !err {
		result.Code = -1
		result.Msg = "Create fail"
		return
	}

	result.Code = 0
	result.Msg = "Create success"
	return
}

func (ass commentSupportService) Add(uid int64, cid int64) (result model.Result) {

	result.Data = commentSupportRepo.GiveSupport(uid, cid)
	if result.Data == false {
		result.Code = -1
		result.Msg = "Add fail"
	}

	result.Code = 0
	result.Msg = "Add success"
	return
}

func (ass commentSupportService) Cancel(uid int64, cid int64) (result model.Result) {
	isExist := commentSupportRepo.IsUserSupported(uid, cid)
	if isExist {
		err := commentSupportRepo.CancelSupport(uid, cid)
		if !err {
			result.Code = -1
			result.Msg = "Cancel fail"
			return
		} else {
			result.Code = 0
			result.Msg = "Cancel Success"
			return
		}
	} else {
		result.Code = -1
		result.Msg = "User has not supported"
		return
	}
}

func (ass commentSupportService) CountSupport(cid int64) int64 {
	return commentSupportRepo.CountSupport(cid)
}

func (ass commentSupportService) HasUserSupported(uid int64, cid int64) bool {
	return commentSupportRepo.IsUserSupported(uid, cid)
}
