package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main1() {
	// 把 123456 换成前端传过来的 MD5 字符串
	password := "e10adc3949ba59abbe56e057f20f883e"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hash))
}
