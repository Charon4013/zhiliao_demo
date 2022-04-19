/**
 * @Author: Pan
 * @Date: 2022/4/12 20:28
 */

package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Hash failed")
		return "", err
	}

	return string(bytes), err
}

func PasswordVerify(password string, hashString string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashString), []byte(password))

	return err == nil
}
