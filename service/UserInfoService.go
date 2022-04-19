/**
 * @Author: Pan
 * @Date: 2022/3/18 21:08
 */

package service

import (
	"fmt"
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/repo"
)

type UserInfoService interface {
	Search(uid int64) (result model.Result)
	Modify(userInfo model.UserInfo) (result model.Result)
	Create(userInfo model.UserInfo) (result model.Result)
}

type userInfoService struct{}

func NewUserInfoService() UserInfoService {
	return &userInfoService{}
}

var userInfoRepo = repo.NewUserInfoRepo()

func (uis userInfoService) Search(uid int64) (result model.Result) {

	userInfoData := userInfoRepo.GetUserInfoByUid(uid)

	result.Code = 0
	result.Data = userInfoData
	return
}

func (uis userInfoService) Modify(userInfo model.UserInfo) (result model.Result) {
	errCode, userInfoData := userInfoRepo.ModifyUserInfo(userInfo)
	if errCode == -1 {
		result.Code = -1
		result.Msg = "Modify fail"
		return
	}

	result.Code = 0
	result.Data = userInfoData
	fmt.Println("result* : ", result)
	return
}

func (uis userInfoService) Create(userInfo model.UserInfo) (result model.Result) {
	errCode, userInfoData := userInfoRepo.CreateUserInfo(userInfo)
	if errCode == -1 {
		result.Code = -1
		result.Msg = "Create fail"
		return
	}
	result.Code = 0
	result.Data = userInfoData
	return
}
