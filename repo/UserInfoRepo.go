/**
 * @Author: Pan
 * @Date: 2022/3/18 20:02
 */

package repo

import (
	"fmt"
	"zhiliao_mvc_demo/datasource"
	"zhiliao_mvc_demo/model"
)

type UserInfoRepo interface {
	// GetUserInfoByUid 查询用户信息
	GetUserInfoByUid(uid int64) model.UserInfo
	// ModifyUserInfo 修改用户信息
	ModifyUserInfo(userInfo model.UserInfo) (int, model.UserInfo)
	// CreateUserInfo 创建用户信息
	CreateUserInfo(userInfo model.UserInfo) (int, model.UserInfo)
}

type userInfoRepo struct{}

func NewUserInfoRepo() UserInfoRepo {
	return &userInfoRepo{}
}

func (uir userInfoRepo) GetUserInfoByUid(uid int64) model.UserInfo {
	var userInfo model.UserInfo

	var userInfoEngine = datasource.Init("UserInfo")
	defer userInfoEngine.Close()

	userInfoEngine.Where("uid = ?", uid).Get(&userInfo)
	fmt.Println("GetUserInfoByUId: uid=", uid)
	return userInfo
}

func (uir userInfoRepo) ModifyUserInfo(userInfo model.UserInfo) (int, model.UserInfo) {

	var userInfoEngine = datasource.Init("UserInfo")
	defer userInfoEngine.Close()

	errCode := 0
	_, err := userInfoEngine.Where("uid = ?", userInfo.Uid).Update(userInfo)
	if err != nil {
		fmt.Println("Update error: ", err)
		errCode = -1
	}
	fmt.Println("userInfo service: ", userInfo)

	return errCode, userInfo
}

func (uir userInfoRepo) CreateUserInfo(userInfo model.UserInfo) (int, model.UserInfo) {

	var userInfoEngine = datasource.Init("UserInfo")
	defer userInfoEngine.Close()

	errCode := 0
	_, err := userInfoEngine.Insert(userInfo)
	if err != nil {
		fmt.Println("Insert error: ", err)
		errCode = -1
	}

	return errCode, userInfo
}
