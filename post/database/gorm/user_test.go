package database_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	database "libai/go/phase-two/post/database/gorm"
	"libai/go/phase-two/post/util"
	"testing"
)

func init() {
	util.InitSlog("../../../log/post.log")
	database.ConnectPostDB("../../conf", "db", util.YAML, "../../../log")
}

func hash(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	digest := hasher.Sum(nil)
	return hex.EncodeToString(digest)
}

func TestRegistUser(t *testing.T) {
	uid, err := database.RegistUser("libai", hash("123456"))
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Printf("注册成功, uid=%d\n", uid)
	}

	uid, err = database.RegistUser("libai", hash("123456"))
	if err == nil {
		fmt.Println("重复注册成功！")
		t.Fail()
	} else {
		fmt.Printf("注册失败: %s\n", err)
	}
}
