/**
 * @Author: Pan
 * @Date: 2022/4/8 18:58
 */

package service

import (
	"fmt"
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/repo"
)

type AnswerService interface {
	Create(a model.Answer) (result model.Result)
	Delete(aid int64) (result model.Result)
	SearchAnswerByAid(qid int64, aid int64) (result model.Result)
	SearchAllUserAnswerByUid(uid int64) (result model.Result)
	SearchAllAnswerByQid(qid int64) (result model.Result)
	Modify(a model.Answer) (result model.Result)
	SearchQuestionAndAnswerListByRandom(num int) (result model.Result)
}

type answerService struct{}

func NewAnswerService() AnswerService {
	return &answerService{}
}

var answerRepo = repo.NewAnswerRepo()

func (as answerService) Create(a model.Answer) (result model.Result) {
	errCode, data := answerRepo.CreateAnswer(a)

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

func (as answerService) Delete(aid int64) (result model.Result) {
	errCode := answerRepo.DeleteAnswer(aid)
	if !errCode {
		result.Code = -1
		result.Msg = "Delete fail"
		return
	}
	result.Code = 0
	result.Msg = "Delete success"
	return
}

func (as answerService) SearchAnswerByAid(qid int64, aid int64) (result model.Result) {
	errCode, data := answerRepo.GetAnswerByAid(qid, aid)
	//fmt.Println("data: ", data)
	if errCode != 0 || data.Id == 0 {
		result.Code = errCode
		result.Msg = "SearchAnswerByAid fail"
		return
	}

	result.Code = 0
	result.Msg = "SearchAnswerByAid success"
	result.Data = data
	return
}

func (as answerService) SearchAllUserAnswerByUid(uid int64) (result model.Result) {
	var questionRepo = repo.NewQuestionRepo()

	errCode, answerData := answerRepo.GetUserAnswerByUid(uid)
	if errCode != 0 || len(answerData) == 0 {
		result.Code = errCode
		result.Msg = "SearchAllUserAnswerByUid fail or len = 0"
		return
	}

	var qaDataList []model.QuestionAnswerData

	for i := 0; i < len(answerData); i++ {
		errCode, questionData := questionRepo.GetQuestionByQid(answerData[i].Qid)

		if errCode == 0 && questionData.Id != 0 {
			qaData := model.QuestionAnswerData{
				Question: questionData,
				Answer:   answerData[i],
			}
			qaDataList = append(qaDataList, qaData)
		}
	}

	result.Code = 0
	result.Msg = "SearchAllUserAnswerByUid success"
	result.Data = qaDataList
	return
}

func (as answerService) SearchAllAnswerByQid(qid int64) (result model.Result) {
	errCode, data := answerRepo.GetAllAnswerByQid(qid)
	if errCode != 0 || data == nil {
		result.Code = errCode
		result.Msg = "SearchAllAnswerByQid fail"
		return
	}
	fmt.Println("service data:", data, len(data))
	result.Code = 0
	result.Msg = "SearchAllAnswerByQid success"
	result.Data = data
	return

}

func (as answerService) Modify(a model.Answer) (result model.Result) {
	errCode, data := answerRepo.ModifyAnswer(a)
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

func (as answerService) SearchQuestionAndAnswerListByRandom(num int) (result model.Result) {
	var questionRepo = repo.NewQuestionRepo()
	var answerSupportRepo = repo.NewAnswerSupportRepo()
	var commentRepo = repo.NewCommentRepo()

	errCode, answerData := answerRepo.GetAnswerListByRandom(num)
	if errCode != 0 {
		result.Code = -1
		result.Msg = "SearchAnswerByRandom fail"
		return
	}
	if len(answerData) == 0 {
		result.Code = -1
		result.Msg = "SearchAnswerByRandom No result"
		return
	}

	var qaDataList []model.QuestionAnswerSupportCommentCount

	for i := 0; i < len(answerData); i++ {

		supportCount := answerSupportRepo.CountSupport(answerData[i].Id)
		commentCount := commentRepo.CountCommentByAid(answerData[i].Id)
		errCode, questionData := questionRepo.GetQuestionByQid(answerData[i].Qid)

		if errCode == 0 && questionData.Id != 0 {
			qaData := model.QuestionAnswerSupportCommentCount{
				Question:     questionData,
				Answer:       answerData[i],
				SupportCount: supportCount,
				CommentCount: commentCount,
			}
			qaDataList = append(qaDataList, qaData)
		}
	}

	result.Code = 0
	result.Msg = "SearchAnswerByRandom success"
	result.Data = qaDataList
	return

}
