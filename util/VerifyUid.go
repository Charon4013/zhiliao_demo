/**
 * @Author: Pan
 * @Date: 2022/4/9 12:32
 */

package util

import "fmt"

func VerifyUid(uid int64, token string) bool {
	tokenUid := GetUidFromParseToken(token)
	fmt.Println("uid: ", uid, " &&tokenUid: ", tokenUid)
	if tokenUid == uid {
		fmt.Println("Uid verify success")
		return true
	}
	fmt.Println("Uid verify fail")
	return false
}
