/**
 * @Author: Pan
 * @Date: 2022/3/19 19:19
 */

package util

import (
	"fmt"
	"strconv"
)

func StringToInt64(str string) (num int64) {
	//fmt.Println("str: ", str)
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("StringToInt64 fail")
	}
	return num
}

func Int64ToString(num int64) (str string) {
	str = strconv.FormatInt(num, 10)
	fmt.Println(str)
	return str
}
