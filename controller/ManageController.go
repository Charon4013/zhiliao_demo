/**
 * @Author: Pan
 * @Date: 2022/4/14 21:03
 */

package controller

import (
	"github.com/kataras/iris/v12"
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/service"
	"zhiliao_mvc_demo/util"
)

type ManageController struct {
	Ctx iris.Context
	Ms  service.ManageService
}

func NewManageController() *ManageController {
	return &ManageController{
		Ms: service.NewManageService(),
	}
}

func (mc ManageController) verifyManageIdentity() model.Result {
	token := mc.Ctx.GetHeader("Token")
	uid := util.GetUidFromParseToken(token)
	isManage := mc.Ms.CheckUserManageIdentityBy(uid)
	if isManage {
		return model.Result{
			Code: 0,
		}
	} else {
		return model.Result{
			Code: -1,
			Msg:  "Check you identity!",
		}
	}
}

func (mc ManageController) PostUser() (result model.Result) {
	result = mc.verifyManageIdentity()
	if result.Code != 0 {
		return result
	}
	return mc.Ms.GetAllUser()
}

func (mc ManageController) PostUserDeleteBy(uid int64) (result model.Result) {
	result = mc.verifyManageIdentity()
	if result.Code != 0 {
		return result
	}
	res := mc.Ms.DeleteUserByUid(uid)

	return res
}

func (mc ManageController) PostUserinfo() (result model.Result) {
	result = mc.verifyManageIdentity()
	if result.Code != 0 {
		return result
	}
	return mc.Ms.GetAllUserInfo()
}

func (mc ManageController) PostUserinfoDeleteBy(uid int64) (result model.Result) {
	result = mc.verifyManageIdentity()
	if result.Code != 0 {
		return result
	}
	res := mc.Ms.DeleteUserInfoByUid(uid)

	return res
}

func (mc ManageController) PostQuestion() (result model.Result) {
	result = mc.verifyManageIdentity()
	if result.Code != 0 {
		return result
	}
	return mc.Ms.GetAllQuestion()
}

func (mc ManageController) PostQuestionDeleteBy(qid int64) (result model.Result) {
	result = mc.verifyManageIdentity()
	if result.Code != 0 {
		return result
	}
	res := mc.Ms.DeleteQuestionByQid(qid)

	return res
}

func (mc ManageController) PostAnswer() (result model.Result) {
	result = mc.verifyManageIdentity()
	if result.Code != 0 {
		return result
	}
	return mc.Ms.GetAllAnswer()
}

func (mc ManageController) PostAnswerDeleteBy(aid int64) (result model.Result) {
	result = mc.verifyManageIdentity()
	if result.Code != 0 {
		return result
	}
	res := mc.Ms.DeleteAnswerByAid(aid)

	return res
}

func (mc ManageController) PostComment() (result model.Result) {
	result = mc.verifyManageIdentity()
	if result.Code != 0 {
		return result
	}
	return mc.Ms.GetAllComment()
}

func (mc ManageController) PostCommentDeleteBy(cid int64) (result model.Result) {
	result = mc.verifyManageIdentity()
	if result.Code != 0 {
		return result
	}
	res := mc.Ms.DeleteCommentByCid(cid)

	return res
}
