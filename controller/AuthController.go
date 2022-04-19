/**
 * @Author: Pan
 * @Date: 2022/4/8 21:38
 */

package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"zhiliao_mvc_demo/middleware"
)

type AuthController struct {
	Ctx iris.Context
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac *AuthController) Get() string {
	str := ac.Ctx.GetHeader("Token")
	fmt.Println("Token: ", str)
	return str
}

func (ac *AuthController) GetDecode() int64 {
	str := ac.Ctx.GetHeader("Token")
	if str == "" {
		return 0
	}
	uid := middleware.GetTokenUid(str)
	if uid == -1 {
		return -1
	}
	fmt.Println(uid)
	return uid
}
