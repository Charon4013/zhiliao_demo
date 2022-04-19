/**
 * @Author: Pan
 * @Date: 2022/3/18 14:02
 */

package service

import (
	"fmt"
	"zhiliao_mvc_demo/middleware"
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/repo"
	"zhiliao_mvc_demo/util"
)

type UserService interface {
	Login(user model.User) (result model.Result)
	Modify(user model.User) (result model.Result)
	Create(user model.User) (result model.Result)
	GetByName(username string) bool
	GetByEmail(email string) bool
	GetByUid(uid int64) bool
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

var userRepo = repo.NewUserRepo()

const secret = "This is my zhiliao secret"

func (us userService) Login(user model.User) (result model.Result) {
	fmt.Println("user: ", user)
	if user.Username == "" {
		result.Code = -1
		result.Msg = "Please input username"
		return
	}
	if user.Password == "" {
		result.Code = -1
		result.Msg = "Please input password"
		return
	}

	pwdHash := userRepo.GetUserPasswordHashByUsername(user.Username)
	fmt.Println("user.Password, pwdHash: ", user.Password, pwdHash)
	verifyResult := util.PasswordVerify(user.Password, pwdHash)
	fmt.Println("verifyResult: ", verifyResult)
	if !verifyResult {
		result.Code = -1
		result.Msg = "Username or password wrong"
		return
	} else {
		user = userRepo.GetUserByUsername(user.Username)
	}

	// token
	token, err := middleware.CreateToken(user.Id)
	if err != nil {
		result.Code = -1
		result.Msg = "Create Token fail"
		return
	}
	user.Token = token

	result.Msg = "Login success"
	result.Data = user
	return
}

// Modify 特指修改密码, 用户名和邮箱不能改
func (us userService) Modify(user model.User) (result model.Result) {

	errCode, userdata := userRepo.ModifyUser(user)
	if errCode == -1 {
		result.Code = -1
		result.Msg = "Modify fail"
		return
	}
	result.Code = 0
	result.Data = userdata
	return
}

func (us userService) Create(user model.User) (result model.Result) {

	if user.Username == "" {
		result.Code = -1
		result.Msg = "Please input username"
		return
	}
	if user.Password == "" {
		result.Code = -1
		result.Msg = "Please input password"
		return
	}

	if us.GetByName(user.Username) && us.GetByEmail(user.Email) {
		result.Code = -1
		result.Msg = "User is exist"
		return
	}

	hash, err := util.PasswordHash(user.Password)
	fmt.Println("hash: ", hash)
	if err != nil {
		result.Code = -1
		result.Msg = "Hash fail"
		return
	}
	user.Password = hash

	errCode, userdata := userRepo.CreateUser(user)
	if errCode == -1 {
		result.Code = -1
		result.Msg = "Create fail"
		return
	}

	// Token
	token, err := middleware.CreateToken(userdata.Id)
	if err != nil {
		result.Code = -1
		result.Msg = "Create Token fail"
		return
	}
	userdata.Token = token

	result.Code = 0
	result.Msg = "Create Success"
	result.Data = userdata
	return
}

func (us userService) GetByName(username string) bool {
	isUsernameExist := userRepo.GetUserByUsername(username)
	if isUsernameExist.Username != "" {
		fmt.Println("User exist")
		return true
	}
	return false
}

func (us userService) GetByEmail(email string) bool {
	isEmailExist := userRepo.GetUserByEmail(email)
	if isEmailExist.Email != "" {
		fmt.Println("User exist")
		return true
	}
	return false
}

func (us userService) GetByUid(uid int64) bool {
	isEmailExist := userRepo.GetUserById(uid)
	if isEmailExist.Email != "" {
		fmt.Println("User exist")
		return true
	}
	return false
}
