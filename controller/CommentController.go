/**
 * @Author: Pan
 * @Date: 2022/4/12 15:22
 */

package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/service"
	"zhiliao_mvc_demo/util"
)

type CommentController struct {
	Ctx iris.Context
	Cs  service.CommentService
	Css service.CommentSupportService
	Ass service.AnswerSupportService
}

func NewCommentController() *CommentController {
	return &CommentController{
		Cs:  service.NewCommentService(),
		Css: service.NewCommentSupportService(),
		Ass: service.NewAnswerSupportService(),
	}
}

func (cc *CommentController) getAid() int64 {
	id := cc.Ctx.Params().Get("aid")
	aid := util.StringToInt64(id)
	return aid
}

func (cc *CommentController) Get() (result model.Result) {
	aid := cc.getAid()
	fmt.Println("aid: ", aid)

	result = cc.Cs.SearchAllAnswerCommentByAid(aid)

	if result.Code != 0 {
		return model.Result{
			Code: -1,
			Msg:  "search comment fail, maybe no comment",
		}
	}
	fmt.Println("result.data: ", result.Data)
	return
}

func (cc *CommentController) PostNew() (result model.Result) {
	var tmp model.CommentStringAidAndUid
	err := cc.Ctx.ReadJSON(&tmp)
	if err != nil {
		fmt.Println("PostNew() ReadJson Error: ", err)
	}

	uid := util.StringToInt64(tmp.Uid)
	aid := util.StringToInt64(tmp.Aid)
	fmt.Println("uid & qid: ", uid, aid)

	c := model.Comment{
		Aid:     aid,
		Uid:     uid,
		Content: tmp.Content,
	}

	token := cc.Ctx.GetHeader("Token")
	fmt.Println("Token: ", token)
	verify := util.VerifyUid(c.Uid, token)
	if !verify {
		result.Msg = "Verify user identity fail"
		result.Code = -1
		return
	}

	result = cc.Cs.Create(c)

	if result.Code != 0 {
		result.Msg = "Comment create fail"
		result.Data = nil
		return
	}

	result.Msg = "NewComment success"
	result.Data = result.Data.(model.Comment)
	fmt.Println("result.Data.(model.Comment): ", result.Data)
	return result
}

func (cc *CommentController) PostSupport() (result model.Result) {
	var tmp model.CommentSupportStringUid
	err := cc.Ctx.ReadJSON(&tmp)
	if err != nil {
		fmt.Println("PostNew() ReadJson Error: ", err)
	}
	fmt.Println("tmp: ", tmp)

	uid := util.StringToInt64(tmp.Uid)
	cid := util.StringToInt64(tmp.Cid)
	fmt.Println("typeof uid & cid: ", uid, cid)

	token := cc.Ctx.GetHeader("Token")
	fmt.Println("Token: ", token)
	verify := util.VerifyUid(uid, token)
	if !verify {
		result.Msg = "Verify user identity fail"
		result.Code = -1
		return
	}

	supportObjIsExist := cc.Css.HasUserSupported(uid, cid)
	if supportObjIsExist {
		result = cc.Css.Add(uid, cid)
		fmt.Println("user has supported before")
	} else {
		result = cc.Css.Create(uid, cid)
	}

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

func (cc *CommentController) PostUnsupport() (result model.Result) {
	var tmp model.CommentSupportStringUid
	err := cc.Ctx.ReadJSON(&tmp)
	if err != nil {
		fmt.Println("PostNew() ReadJson Error: ", err)
	}

	uid := util.StringToInt64(tmp.Uid)
	cid := util.StringToInt64(tmp.Cid)
	fmt.Println("typeof uid & cid: ", uid, cid)

	token := cc.Ctx.GetHeader("Token")
	fmt.Println("Token: ", token)
	verify := util.VerifyUid(uid, token)
	if !verify {
		result.Msg = "Verify user identity fail"
		result.Code = -1
		return
	}

	supportObjIsExist := cc.Css.HasUserSupported(uid, cid)
	if supportObjIsExist {
		fmt.Println("user has supported before")

		result = cc.Css.Cancel(uid, cid)
		if result.Code != 0 {
			return model.Result{
				Code: -1,
				Msg:  "Unsupport fail",
			}
		} else {
			return model.Result{
				Code: 0,
				Msg:  "Unsupport success",
			}
		}
	} else {
		return model.Result{
			Code: -1,
			Msg:  "User has not support before",
		}
	}
}

func (cc *CommentController) GetSupportBy(cid int64) (result model.Result) {

	type resultData struct {
		Count     int64
		Supported bool
	}

	isUserSupported := false

	fmt.Println("GetSupport")
	//id := cc.Ctx.Params().Get("cid")
	//cid := util.StringToInt64(id)
	var token string
	token = cc.Ctx.GetHeader("Token")
	fmt.Println("GetHeader token:", string(token))
	fmt.Println("1233:", token)
	if token == "null" || token == "" || token == " null" {
		fmt.Println("Not Login")
		return model.Result{
			Code: -1,
			Msg:  "Not Login",
		}
	} else {
		uid := util.GetUidFromParseToken(token)
		if uid == 0 {
			fmt.Println("Token Error")
			isUserSupported = false
		} else {
			isUserSupported = cc.Css.HasUserSupported(uid, cid)
		}
	}
	count := cc.Css.CountSupport(cid)
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

func (cc *CommentController) GetCount() (result model.Result) {
	fmt.Println("GetCount")
	aid := cc.getAid()
	fmt.Println("aid: ", aid)
	return cc.Cs.GetCommentCountByAid(aid)
}
