/**
 * @Author: Pan
 * @Date: 2022/3/28 12:19
 */

package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/service"
	"zhiliao_mvc_demo/util"
)

type QuestionController struct {
	Ctx iris.Context
	Qs  service.QuestionService
	As  service.AnswerService
}

func NewQuestionController() *QuestionController {
	return &QuestionController{Qs: service.NewQuestionService()}
}

func (qc *QuestionController) PostNew() model.Result {

	// 前端localstorage存的uid是字符串，要转换，不然ReadJSON报错
	var tmp model.QuestionStringUid
	err := qc.Ctx.ReadJSON(&tmp)
	if err != nil {
		fmt.Println("PostNew() ReadJson Error: ", err)
	}

	uid := util.StringToInt64(tmp.Uid)

	q := model.Question{
		Uid:         uid,
		Title:       tmp.Title,
		Description: tmp.Description,
	}

	result := qc.Qs.Create(q)

	//fmt.Println("result.Code:", result.Code)
	return result
}

func (qc *QuestionController) GetRandom() model.Result {
	num := 10
	//fmt.Println(qc.Qs.SearchQuestionListByRandom(num))
	return qc.Qs.SearchQuestionListByRandom(num)
}

func (qc *QuestionController) GetSearchBy(str string) model.Result {

	//err := qc.Ctx.ReadJSON(&str)
	//if err != nil {
	//	return model.Result{
	//		Code: -1,
	//		Msg:  "GetQuestionByString fail",
	//		Data: nil,
	//	}
	//}
	result := qc.Qs.Search(str)
	return result
}

func (qc *QuestionController) PostModify() model.Result {

	// 前端localstorage存的uid是字符串，要转换，不然ReadJSON报错
	var tmp model.QuestionStringUid
	err := qc.Ctx.ReadJSON(&tmp)
	if err != nil {
		fmt.Println("PostNew() ReadJson Error: ", err)
	}

	uid := util.StringToInt64(tmp.Uid)

	q := model.Question{
		Uid:         uid,
		Title:       tmp.Title,
		Description: tmp.Description,
	}

	result := qc.Qs.Modify(q)

	//fmt.Println("result.Code:", result.Code)
	return result
}

func (qc *QuestionController) GetBy() {

}
