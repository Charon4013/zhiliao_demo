/**
 * @Author: Pan
 * @Date: 2022/3/18 19:20
 */

package util

import (
	"fmt"
	"math/rand"
	"time"
)

// CreateRandomString 生成10位随机字符串
func CreateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	randomString := randomString(10)
	fmt.Println("randomString:", randomString)
	return randomString
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(97, 122))
	}
	return string(bytes)
}
