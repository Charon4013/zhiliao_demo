/**
 * @Author: Pan
 * @Date: 2022/4/14 21:43
 */

package service

import (
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/repo"
)

type ManageService interface {
	GetAllUser() (result model.Result)
	DeleteUserByUid(uid int64) (result model.Result)
	GetAllUserInfo() (result model.Result)
	DeleteUserInfoByUid(uid int64) (result model.Result)
	GetAllQuestion() (result model.Result)
	DeleteQuestionByQid(qid int64) (result model.Result)
	GetAllAnswer() (result model.Result)
	DeleteAnswerByAid(aid int64) (result model.Result)
	GetAllComment() (result model.Result)
	DeleteCommentByCid(cid int64) (result model.Result)
	CheckUserManageIdentityBy(uid int64) bool
}

type manageService struct{}

func NewManageService() ManageService {
	return &manageService{}
}

var manageRepo = repo.NewManageRepo()

func (ms manageService) GetAllUser() (result model.Result) {
	errCode, data := manageRepo.GetAllUser()
	if errCode != 0 {
		result.Code = errCode
		result.Msg = "GetAllUser fail"
		return
	}
	result.Code = 0
	result.Msg = "GetAllUser success"
	result.Data = data
	return
}

func (ms manageService) DeleteUserByUid(uid int64) (result model.Result) {
	err := manageRepo.DeleteUserByUid(uid)
	if !err {
		result.Code = -1
		result.Msg = "DeleteUserByUid fail"
		return
	}
	result.Code = 0
	result.Msg = "DeleteUserByUid success"
	return
}

func (ms manageService) GetAllUserInfo() (result model.Result) {
	errCode, data := manageRepo.GetAllUserInfo()
	if errCode != 0 {
		result.Code = errCode
		result.Msg = "GetAllUserInfo fail"
		return
	}
	result.Code = 0
	result.Msg = "GetAllUserInfo success"
	result.Data = data
	return
}

func (ms manageService) DeleteUserInfoByUid(uid int64) (result model.Result) {
	err := manageRepo.DeleteUserInfoById(uid)
	if !err {
		result.Code = -1
		result.Msg = "DeleteUserInfoById fail"
		return
	}
	result.Code = 0
	result.Msg = "DeleteUserInfoById success"
	return
}

func (ms manageService) GetAllQuestion() (result model.Result) {
	errCode, data := manageRepo.GetAllQuestion()
	if errCode != 0 {
		result.Code = errCode
		result.Msg = "GetAllQuestion fail"
		return
	}
	result.Code = 0
	result.Msg = "GetAllQuestion success"
	result.Data = data
	return
}

func (ms manageService) DeleteQuestionByQid(qid int64) (result model.Result) {
	err := manageRepo.DeleteQuestionByQid(qid)
	if !err {
		result.Code = -1
		result.Msg = "DeleteQuestionByQid fail"
		return
	}
	result.Code = 0
	result.Msg = "DeleteQuestionByQid success"
	return
}

func (ms manageService) GetAllAnswer() (result model.Result) {
	errCode, data := manageRepo.GetAllAnswer()
	if errCode != 0 {
		result.Code = errCode
		result.Msg = "GetAllAnswer fail"
		return
	}
	result.Code = 0
	result.Msg = "GetAllAnswer success"
	result.Data = data
	return
}

func (ms manageService) DeleteAnswerByAid(aid int64) (result model.Result) {
	err := manageRepo.DeleteAnswerByAid(aid)
	if !err {
		result.Code = -1
		result.Msg = "DeleteAnswerByAid fail"
		return
	}
	result.Code = 0
	result.Msg = "DeleteAnswerByAid success"
	return
}

func (ms manageService) GetAllComment() (result model.Result) {
	errCode, data := manageRepo.GetAllComment()
	if errCode != 0 {
		result.Code = errCode
		result.Msg = "GetAllComment fail"
		return
	}
	result.Code = 0
	result.Msg = "GetAllComment success"
	result.Data = data
	return
}

func (ms manageService) DeleteCommentByCid(cid int64) (result model.Result) {
	err := manageRepo.DeleteCommentByCid(cid)
	if !err {
		result.Code = -1
		result.Msg = "DeleteCommentByCid fail"
		return
	}
	result.Code = 0
	result.Msg = "DeleteCommentByCid success"
	return
}

func (ms manageService) CheckUserManageIdentityBy(uid int64) bool {
	return manageRepo.CheckUserManageIdentityBy(uid)
}
