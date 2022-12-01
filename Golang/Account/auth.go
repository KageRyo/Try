package main

import (
	"errors"
)

var UserData map[string]string

func init() {
	UserData = map[string]string{
		"admin": "password",
	}
}

func CheckUserIsExist(username string) bool {
	_, isExist := UserData[username]
	return isExist
}

func CheckPassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	} else {
		return errors.New("帳號或密碼錯誤")
	}
}

func Auth(username string, password string) error {
	if isExist := CheckUserIsExist(username); isExist {
		return CheckPassword(UserData[username], password)
	} else {
		return errors.New("使用者不存在")
	}
}
