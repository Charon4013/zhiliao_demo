/**
 * @Author: Pan
 * @Date: 2022/4/10 17:04
 */

package service

import (
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/repo"
)

type AnswerSupportService interface {
	Create(uid int64, aid int64) (result model.Result)
	Add(uid int64, aid int64) (result model.Result)
	Cancel(uid int64, aid int64) (result model.Result)
	CountSupport(aid int64) int64
	HasUserSupported(uid int64, aid int64) bool
}

type answerSupportService struct{}

func NewAnswerSupportService() AnswerSupportService {
	return &answerSupportService{}
}

var answerSupportRepo = repo.NewAnswerSupportRepo()

func (ass answerSupportService) Create(uid int64, aid int64) (result model.Result) {
	err := answerSupportRepo.CreateSupport(uid, aid)

	if !err {
		result.Code = -1
		result.Msg = "Create fail"
		return
	}

	result.Code = 0
	result.Msg = "Create success"
	return
}

func (ass answerSupportService) Add(uid int64, aid int64) (result model.Result) {

	hasRecorded := answerSupportRepo.IsUserRecorded(uid, aid)
	if !hasRecorded {
		ass.Create(uid, aid)
	}

	result.Data = answerSupportRepo.GiveSupport(uid, aid)
	if result.Data == false {
		result.Code = -1
		result.Msg = "Add fail"
	}

	result.Code = 0
	result.Msg = "Add success"
	return
}

func (ass answerSupportService) Cancel(uid int64, aid int64) (result model.Result) {
	hasRecorded := answerSupportRepo.IsUserRecorded(uid, aid)
	if !hasRecorded {
		result.Code = -1
		result.Msg = "Cancel fail, user has not support before"
		return
	}

	err := answerSupportRepo.CancelSupport(uid, aid)
	if !err {
		result.Code = -1
		result.Msg = "Cancel fail"
		return
	} else {
		result.Code = 0
		result.Msg = "Cancel Success"
		return
	}
}

func (ass answerSupportService) CountSupport(aid int64) int64 {
	return answerSupportRepo.CountSupport(aid)
}

func (ass answerSupportService) HasUserSupported(uid int64, aid int64) bool {
	return answerSupportRepo.IsUserSupported(uid, aid)
}
