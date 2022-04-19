/**
 * @Author: Pan
 * @Date: 2022/3/18 15:43
 */

package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/service"
)

type UserController struct {
	Ctx     iris.Context
	Us      service.UserService
	Uis     service.UserInfoService
	Qs      service.QuestionService
	Session *sessions.Session
}

func NewUserController() *UserController {
	return &UserController{
		Us:  service.NewUserService(),
		Uis: service.NewUserInfoService(),
	}
}

// 注册登录： get&post:/user/login | post:/user/register

func (uc *UserController) GetLogin() {
	//uc.Ctx.View("login.html")
}

func (uc *UserController) PostLogin() model.Result {
	//username := uc.Ctx.PostValue("username")
	//password := uc.Ctx.PostValue("password")
	user := model.User{}
	err := uc.Ctx.ReadJSON(&user)
	fmt.Println("#1-UserController# ReadJSON user: ", user)
	if err != nil {
		fmt.Println("#2-UserController# ReadJSON error")
	}
	result := uc.Us.Login(user)
	fmt.Println("#3-UserController# result:", result)

	if result.Code != 0 {
		return model.Result{
			Code: -1,
			Msg:  "Login fail, please check your input",
		}
	}
	user = result.Data.(model.User)
	//fmt.Println(" result.Data.(model.User): --->", result.Data.(model.User).Username)
	//fmt.Println(" result.Data.(model.User): --->", result.Data.(model.User).Manage)
	//fmt.Println(" result.Data.(model.User): --->", result.Data.(model.User).Token)

	userResponse := model.User{
		Id:       user.Id,
		Username: user.Username,
		Manage:   user.Manage,
		Token:    user.Token,
	}

	fmt.Println("#4-UserController# userResponse", userResponse)

	//token, err := middleware.CreateToken(data.)
	//fmt.Println("login user data: ", data)
	return model.Result{
		Code: 0,
		Msg:  "Login Success",
		Data: userResponse,
	}

}

func (uc *UserController) PostRegister() model.Result {
	user := model.User{}
	err := uc.Ctx.ReadJSON(&user)

	fmt.Println("#UserController# Read user: ", user)
	if err != nil {
		fmt.Println("#UserController# ReadJSON error")
	}

	// 判断有没有重名或一个邮箱注册多个号
	if uc.Us.GetByEmail(user.Email) || uc.Us.GetByName(user.Username) {
		return model.Result{
			Code: -1,
			Msg:  "Username or Email has been registered",
		}
	}

	result := uc.Us.Create(user)
	fmt.Println("#UserController# result: ", result)

	errCode := 0
	var msg, msg1, msg2 string
	result2 := model.Result{}

	if result.Code == 0 {
		user = result.Data.(model.User)
		fmt.Println("created user: ", user)
		// 注册一个空的用户信息
		result2 = uc.Uis.Create(model.UserInfo{Uid: user.Id})

		fmt.Println("#UserController# result2: ", result2)
	} else {
		fmt.Println("result.Code != 0")
		msg1 = "Create user fail"
		errCode = -1
	}

	if result2.Code != 0 {
		msg2 = "Create userInfo fail"
		errCode = -1
	}
	if msg1 == "" && msg2 == "" {
		msg = "Register success"
	} else {
		msg = msg1 + " " + msg2
	}

	user = result.Data.(model.User)

	userResponse := model.User{
		Id:       user.Id,
		Username: user.Username,
		Manage:   user.Manage,
		Token:    user.Token,
	}

	fmt.Println("#4-UserController# userResponse", userResponse)

	return model.Result{
		Code: errCode,
		Msg:  msg,
		Data: userResponse,
	}
}
