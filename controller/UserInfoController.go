/**
 * @Author: Pan
 * @Date: 2022/4/7 15:26
 */

package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/service"
	"zhiliao_mvc_demo/util"
)

type UserInfoController struct {
	Ctx iris.Context
	Qs  service.QuestionService
	Uis service.UserInfoService
	As  service.AnswerService
}

func NewUserInfoController() *UserInfoController {
	return &UserInfoController{
		Qs:  service.NewQuestionService(),
		Uis: service.NewUserInfoService(),
		As:  service.NewAnswerService(),
	}
}

func (uic *UserInfoController) getUid() int64 {
	id := uic.Ctx.Params().Get("uid")
	uid := util.StringToInt64(id)
	return uid
}

func (uic *UserInfoController) Get() (result model.Result) {
	uid := uic.getUid()
	//if err := uic.Ctx.ReadJSON(&uid); err != nil {
	//	fmt.Println("PostCreate() ReadJson Error: ", err)
	//	result.Msg = "Data error"
	//}
	fmt.Println(uic.Uis.Search(uid))
	return uic.Uis.Search(uid)
}

func (uic *UserInfoController) Post() (result model.Result) {
	userInfo := model.UserInfo{}
	if err := uic.Ctx.ReadJSON(&userInfo); err != nil {
		fmt.Println("Post() ReadJson Error: ", err)
		result.Msg = "Data error"
		return result
	}

	// 验证token
	token := uic.Ctx.GetHeader("Token")
	fmt.Println("Token: ", token)
	verify := util.VerifyUid(userInfo.Uid, token)
	if !verify {
		result.Msg = "Verify user identity fail"
		result.Code = -1
		return
	}

	result = uic.Uis.Modify(userInfo)
	if result.Code != 0 {
		result.Msg = "No userInfo"
		result.Data = nil
		return
	}

	result.Msg = "Post userInfo success"
	result.Data = result.Data.(model.UserInfo)
	fmt.Println("uc.Uis.Modify(userInfo): ", result.Data)
	return result
}

func (uic *UserInfoController) GetQuestion() model.Result {
	uid := uic.getUid()
	result := uic.Qs.SearchUserAllList(uid)
	return result
}

func (uic *UserInfoController) GetAnswer() model.Result {
	uid := uic.getUid()
	result := uic.As.SearchAllUserAnswerByUid(uid)
	return result
}

func (uic *UserInfoController) PostAnswerDeleteBy(aid int64) model.Result {

	uid := uic.getUid()

	token := uic.Ctx.GetHeader("Token")
	verify := util.VerifyUid(uid, token)
	var result model.Result
	if !verify {
		result.Msg = "Verify user identity fail"
		result.Code = -1
		return result
	}

	return uic.As.Delete(aid)
}

func (uic *UserInfoController) PostQuestionDeleteBy(qid int64) model.Result {

	uid := uic.getUid()

	token := uic.Ctx.GetHeader("Token")
	verify := util.VerifyUid(uid, token)
	var result model.Result
	if !verify {
		result.Msg = "Verify user identity fail"
		result.Code = -1
		return result
	}

	return uic.Qs.Delete(qid)
}

func (uic *UserInfoController) PostQuestionModifyBy(qid int64) model.Result {
	uid := uic.getUid()

	var q model.Question
	err := uic.Ctx.ReadJSON(&q)
	if err != nil {
		fmt.Println("Post() ReadJson Error: ", err)
		return model.Result{
			Code: -1,
			Msg:  "Read JSON fail",
		}
	}

	token := uic.Ctx.GetHeader("Token")
	verify := util.VerifyUid(uid, token)
	var result model.Result
	if !verify {
		result.Msg = "Verify user identity fail"
		result.Code = -1
		return result
	}
	result = uic.Qs.Modify(q)

	return result
}
