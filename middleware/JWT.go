/**
 * @Author: Pan
 * @Date: 2022/3/23 16:17
 */

package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Authorized bool
	Uid        int64
	Iss        string
}

var hmacSampleSecret []byte

var secret = []byte("This is a zhiliao JWT secret!!!")

func CreateToken(uid int64) (string, error) {
	fmt.Println("uid :", uid)
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"uid":        uid,
		"iss":        "Pan",
		//"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	token, err := at.SignedString(secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseToken interface返回的是float64，还是要处理下
//func ParseToken(tokenString string, secret string) (int64, error) {
//	claim, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return []byte(secret), nil
//	})
//	if err != nil {
//		fmt.Println(" err", err)
//		return 0, err
//	}
//	uid := claim.Claims.(jwt.MapClaims)["uid"].(float64)
//	fmt.Println("uid:  ", uid)
//	return int64(uid), nil
//}

func ParseToken(tokenStr string, secret []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		fmt.Println("token, err: ", token, err)
		return secret, nil
	})
	claims = token.Claims
	//fmt.Println("claims: ", claims)
	return
}

func GetTokenUid(tokenStr string) int64 {
	claims, err := ParseToken(tokenStr, secret)
	//fmt.Println("claims: ", claims)
	if err != nil {
		fmt.Println("ParseToken error: ", err)
		return -1
	}
	var uid int64
	id := claims.(jwt.MapClaims)["uid"].(float64)
	//fmt.Println("id: ", id)
	uid = int64(id)
	//fmt.Println("uid:", uid)
	return uid
}
