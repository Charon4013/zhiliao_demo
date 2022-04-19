/**
 * @Author: Pan
 * @Date: 2022/4/8 20:39
 */

package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/service"
	"zhiliao_mvc_demo/util"
)

type AnswerController struct {
	Ctx iris.Context
	Qs  service.QuestionService
	As  service.AnswerService
}

func NewAnswerController() *AnswerController {
	return &AnswerController{
		As: service.NewAnswerService(),
		Qs: service.NewQuestionService(),
	}
}

func (ac *AnswerController) getQid() int64 {
	id := ac.Ctx.Params().Get("qid")
	qid := util.StringToInt64(id)
	return qid
}

func (ac *AnswerController) Get() (result model.Result) {
	qid := ac.getQid()
	//fmt.Println("qi1:", qid, reflect.TypeOf(qid))

	questionResult := ac.Qs.SearchOne(qid)
	if questionResult.Code != 0 {
		fmt.Println("qc err: ", questionResult.Msg)
		return model.Result{
			Code: -1,
			Msg:  "Question not found",
		}
	}
	qadata := model.QuestionAnswerDataList{
		Question: model.Question{},
		Answer:   nil,
	}

	qadata.Question = (questionResult.Data).(model.Question)

	answerResult := ac.As.SearchAllAnswerByQid(qid)
	if answerResult.Code != 0 {
		result.Code = 0
		result.Msg = "Has no answers yet"
		qadata.Answer = nil
	} else {
		qadata.Answer = (answerResult.Data).([]model.Answer)
	}

	fmt.Println("qadata: ", qadata)
	result.Data = qadata

	return result
}

func (ac *AnswerController) GetAnswerBy(aid int64) (result model.Result) {
	qid := ac.getQid()
	//fmt.Println("qid&aid: ", qid, aid)

	questionResult := ac.Qs.SearchOne(qid)
	//fmt.Println("questionResult: ", questionResult)

	if questionResult.Code != 0 {
		fmt.Println("qc err: ", questionResult.Msg)
		return model.Result{
			Code: -1,
			Msg:  "Question not found",
		}
	}

	qadata := model.QuestionAnswerData{}
	qadata.Question = (questionResult.Data).(model.Question)

	answerResult := ac.As.SearchAnswerByAid(qid, aid)
	if answerResult.Code != 0 {
		result.Code = 0
		result.Msg = "Has no answers yet"
		qadata.Answer = model.Answer{}
	} else {
		qadata.Answer = (answerResult.Data).(model.Answer)
	}
	fmt.Println("answerResult: ", answerResult)
	result.Data = qadata
	return result
}

func (ac *AnswerController) GetSingle() (result model.Result) {
	qid := ac.getQid()

	result = ac.Qs.SearchOne(qid)
	fmt.Println("questionResult: ", result)
	if result.Code != 0 {
		fmt.Println("qc err: ", result.Msg)
		return model.Result{
			Code: -1,
			Msg:  "Question not found",
		}
	}

	return
}

func (ac *AnswerController) PostAnswerNew() (result model.Result) {

	var tmp model.AnswerStringUidAndQid
	err := ac.Ctx.ReadJSON(&tmp)
	if err != nil {
		fmt.Println("PostNew() ReadJson Error: ", err)
	}

	uid := util.StringToInt64(tmp.Uid)
	qid := util.StringToInt64(tmp.Qid)
	fmt.Println("uid & qid: ", uid, qid)

	a := model.Answer{
		Qid:     qid,
		Uid:     uid,
		Content: tmp.Content,
	}

	token := ac.Ctx.GetHeader("Token")
	fmt.Println("Token: ", token)
	verify := util.VerifyUid(a.Uid, token)
	if !verify {
		result.Msg = "Verify user identity fail"
		result.Code = -1
		return
	}

	result = ac.As.Create(a)

	if result.Code != 0 {
		result.Msg = "Answer create fail"
		result.Data = nil
		return
	}

	result.Msg = "PostAnswerNew success"
	result.Data = result.Data.(model.Answer)
	fmt.Println("ac.As.Create(answer): ", result.Data)
	return result
}

// PostModify modify Question, only question
func (ac *AnswerController) PostModify() (result model.Result) {
	var tmp model.QuestionStringIdUid
	err := ac.Ctx.ReadJSON(&tmp)
	if err != nil {
		fmt.Println("PostModify() ReadJson Error: ", err)
	}

	uid := util.StringToInt64(tmp.Uid)
	qid := util.StringToInt64(tmp.Id)
	fmt.Println("uid & qid: ", uid, qid)

	q := model.Question{
		Id:          qid,
		Uid:         uid,
		Title:       tmp.Title,
		Description: tmp.Description,
	}

	token := ac.Ctx.GetHeader("Token")
	fmt.Println("Token: ", token)
	verify := util.VerifyUid(q.Uid, token)
	if !verify {
		result.Msg = "Verify user identity fail"
		result.Code = -1
		return
	}

	result = ac.Qs.Modify(q)

	if result.Code != 0 {
		result.Msg = "Question modify fail"
		result.Data = nil
		return
	}

	result.Msg = "PostModify success"
	result.Data = result.Data.(model.Question)
	return result
}
