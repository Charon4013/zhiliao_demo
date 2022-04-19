/**
 * @Author: Pan
 * @Date: 2022/4/14 15:47
 */

package controller

import (
	"github.com/kataras/iris/v12"
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/service"
)

type MainPageController struct {
	Ctx iris.Context
	Qs  service.QuestionService
	As  service.AnswerService
}

func NewMainPageController() *MainPageController {
	return &MainPageController{
		As: service.NewAnswerService(),
		Qs: service.NewQuestionService(),
	}
}

func (mpc MainPageController) GetExplore() (result model.Result) {
	// num不起作用，sql暂时写死了10条
	result = mpc.As.SearchQuestionAndAnswerListByRandom(10)
	return
}
