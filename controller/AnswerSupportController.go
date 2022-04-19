/**
 * @Author: Pan
 * @Date: 2022/4/11 13:43
 */

package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/service"
	"zhiliao_mvc_demo/util"
)

type AnswerSupportController struct {
	Ctx iris.Context
	Ss  service.AnswerSupportService
	As  service.AnswerService
}

func NewAnswerSupportController() *AnswerSupportController {
	return &AnswerSupportController{
		Ss: service.NewAnswerSupportService(),
		As: service.NewAnswerService(),
	}
}

func (asc AnswerSupportController) PostSupport() (result model.Result) {
	var tmp model.AnswerSupportStringUid
	err := asc.Ctx.ReadJSON(&tmp)
	if err != nil {
		fmt.Println("PostNew() ReadJson Error: ", err)
	}

	uid := util.StringToInt64(tmp.Uid)
	aid := util.StringToInt64(tmp.Aid)
	fmt.Println("typeof uid & aid: ", uid, aid)

	token := asc.Ctx.GetHeader("Token")
	fmt.Println("Token: ", token)
	verify := util.VerifyUid(uid, token)
	if !verify {
		result.Msg = "Verify user identity fail"
		result.Code = -1
		return
	}

	result = asc.Ss.Add(uid, aid)

	if result.Code != 0 {
		return model.Result{
			Code: -1,
			Msg:  "Support fail",
		}
	}
	return model.Result{
		Code: 0,
		Msg:  "Support success",
	}
}

func (asc AnswerSupportController) PostUnsupport() (result model.Result) {
	var tmp model.AnswerSupportStringUid
	err := asc.Ctx.ReadJSON(&tmp)
	if err != nil {
		fmt.Println("PostNew() ReadJson Error: ", err)
	}

	uid := util.StringToInt64(tmp.Uid)
	aid := util.StringToInt64(tmp.Aid)
	fmt.Println("typeof uid & aid: ", uid, aid)

	token := asc.Ctx.GetHeader("Token")
	fmt.Println("Token: ", token)
	verify := util.VerifyUid(uid, token)
	if !verify {
		result.Msg = "Verify user identity fail"
		result.Code = -1
		return
	}

	result = asc.Ss.Cancel(uid, aid)
	return
}

func (asc AnswerSupportController) GetSupport() (result model.Result) {

	type resultData struct {
		Count     int64
		Supported bool
	}

	isUserSupported := false

	fmt.Println("GetSupport AnswerSupportController")

	id := asc.Ctx.Params().Get("aid")
	aid := util.StringToInt64(id)

	token := asc.Ctx.GetHeader("Token")
	//fmt.Println("GetHeader token AnswerSupportController: ", token)
	//fmt.Println("typeof", reflect.TypeOf(token))
	var uid int64
	if token == "" || token == "null" || token == "No token" {
		fmt.Println("Not Login, token:", token)

		result.Code = -1
		result.Msg = "Not Login"

	} else {
		uid = util.GetUidFromParseToken(token)
	}
	if uid <= 0 {
		fmt.Println("Token Error")
		isUserSupported = false
	} else {
		isUserSupported = asc.Ss.HasUserSupported(uid, aid)
	}

	count := asc.Ss.CountSupport(aid)
	resultdata := resultData{
		Count:     count,
		Supported: isUserSupported,
	}

	fmt.Println("GetSupport() result: ", count, isUserSupported)
	return model.Result{
		Code: 0,
		Msg:  "GetSupport success",
		Data: resultdata,
	}
}

func (asc AnswerSupportController) PostModify() (result model.Result) {
	var tmp model.AnswerStringIdUidQid
	err := asc.Ctx.ReadJSON(&tmp)
	if err != nil {
		fmt.Println("PostModify() ReadJson Error: ", err)
	}

	aid := util.StringToInt64(tmp.Id)
	uid := util.StringToInt64(tmp.Uid)

	a := model.Answer{
		Id:      aid,
		Uid:     uid,
		Content: tmp.Content,
	}

	token := asc.Ctx.GetHeader("Token")
	fmt.Println("Token: ", token)
	verify := util.VerifyUid(a.Uid, token)
	if !verify {
		result.Msg = "Verify user identity fail"
		result.Code = -1
		return
	}

	result = asc.As.Modify(a)

	if result.Code != 0 {
		result.Msg = "Answer modify fail"
		result.Data = nil
		return
	}

	result.Msg = "PostModify success"
	result.Data = result.Data.(model.Answer)

	return
}
