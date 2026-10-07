package xorm

import (
	"fmt"
	"log/slog"
	"math/rand/v2"

	"xorm.io/xorm"
)

func Create(engine *xorm.Engine) {
	user := User{
		UserId:   rand.IntN(100000),
		Degree:   "本科",
		Gender:   "男",
		City:     "上海",
		Keywords: []string{"编程", "golang"},
	}
	affected, err := engine.Insert(&user)
	if err != nil {
		slog.Error("插入记录失败", "error", err)
	}
	fmt.Printf("after insert user is %#v\n", user)
	fmt.Printf("影响%d行\n", affected)

	user1 := user
	user1.Id = 0
	user1.UserId = rand.IntN(100000)
	user2 := user
	user2.Id = 0
	user2.UserId = rand.IntN(100000)
	users := []User{user1, user2}
	affected, err = engine.Insert(users)
	if err != nil {
		slog.Error("插入记录失败", "error", err)
	}
	fmt.Printf("影响%d行\n", affected)
}
