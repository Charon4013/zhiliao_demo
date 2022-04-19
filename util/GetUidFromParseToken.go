/**
 * @Author: Pan
 * @Date: 2022/4/9 12:29
 */

package util

import (
	"zhiliao_mvc_demo/middleware"
)

func GetUidFromParseToken(token string) int64 {
	if token == "" {
		return 0
	}
	uid := middleware.GetTokenUid(token)
	if uid == -1 {
		return -1
	}
	//fmt.Println(uid)
	return uid
}
