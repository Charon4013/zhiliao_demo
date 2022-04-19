/**
 * @Author: Pan
 * @Date: 2022/3/28 12:04
 */

package service

import (
	"fmt"
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/repo"
)

type QuestionService interface {
	SearchOne(qid int64) (result model.Result)
	Search(str string) (result model.Result)
	Create(q model.Question) (result model.Result)
	Modify(q model.Question) (result model.Result)
	Delete(qid int64) (result model.Result)
	SearchUserAllList(uid int64) (result model.Result)
	SearchQuestionListByRandom(num int) (result model.Result)
}

type questionService struct{}

func NewQuestionService() QuestionService {
	return &questionService{}
}

var questionRepo = repo.NewQuestionRepo()

func (qs questionService) SearchOne(qid int64) (result model.Result) {
	fmt.Println("SearchOne=========================")
	errCode, data := questionRepo.GetQuestionByQid(qid)
	//fmt.Println("code&data",errCode, data)
	if errCode != 0 || data.Id == 0 {
		result.Code = errCode
		result.Msg = "SearchOne fail"
		return
	}
	//fmt.Println("service data:", data, data.Id)
	result.Code = 0
	result.Msg = "SearchOne success"
	result.Data = data
	fmt.Println("SearchOne=========================END")
	return
}

func (qs questionService) Search(str string) (result model.Result) {
	errCode, data := questionRepo.GetQuestionListByTitle(str)
	if errCode != 0 || len(data) == 0 {
		result.Code = errCode
		result.Msg = "Search fail"
		return
	}
	fmt.Println("service data:", data, len(data))
	result.Code = 0
	result.Msg = "Search success"
	result.Data = data
	return
}

func (qs questionService) Create(q model.Question) (result model.Result) {
	errCode, data := questionRepo.CreateQuestion(q)

	if errCode != 0 {
		result.Code = errCode
		result.Msg = "Create fail"
		return
	}
	result.Code = 0
	result.Msg = "Create success"
	result.Data = data
	return
}

func (qs questionService) Modify(q model.Question) (result model.Result) {
	errCode, data := questionRepo.ModifyQuestion(q)
	if errCode != 0 {
		result.Code = errCode
		result.Msg = "Modify fail"
		return
	}
	result.Code = 0
	result.Msg = "Modify success"
	result.Data = data
	return
}

func (qs questionService) Delete(qid int64) (result model.Result) {
	errCode := questionRepo.DeleteQuestion(qid)
	if !errCode {
		result.Code = -1
		result.Msg = "Delete fail"
		return
	}
	result.Code = 0
	result.Msg = "Delete success"
	return
}

func (qs questionService) SearchUserAllList(uid int64) (result model.Result) {
	errCode, data := questionRepo.GetQuestionListByUid(uid)
	if errCode != 0 {
		result.Code = -1
		result.Msg = "SearchUserAllList fail"
		return
	}
	if len(data) == 0 {
		result.Code = -1
		result.Msg = "SearchUserAllList No result"
		return
	}
	result.Code = 0
	result.Msg = "SearchUserAllList success"
	result.Data = data
	return
}

func (qs questionService) SearchQuestionListByRandom(num int) (result model.Result) {
	errCode, data := questionRepo.GetQuestionListByRandom(num)
	if errCode != 0 {
		result.Code = -1
		result.Msg = "SearchRandomList fail"
		return
	}
	if len(data) == 0 {
		result.Code = -1
		result.Msg = "SearchRandomList No result"
		return
	}
	result.Code = 0
	result.Msg = "SearchRandomList success"
	result.Data = data
	return
}
