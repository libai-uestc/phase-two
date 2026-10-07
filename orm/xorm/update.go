package xorm

import (
	"fmt"
	"time"

	"xorm.io/xorm"
)

func Update(engine *xorm.Engine) {
	affected, _ := engine.Table(User{}).Where("city=?", "上海").Update(map[string]any{"degree": "硕士", "gender": "男", "keywords": ""})
	fmt.Printf("更新了%d行\n", affected)

	affected, _ = engine.Where("city=?", "上海").MustCols("keywords").Update(User{Degree: "本科", Gender: "男"})
	fmt.Printf("更新了%d行\n", affected)
}

func UpdateByVersion(engine *xorm.Engine) {
	var user User
	ok, err := engine.ID(8).Get(&user)
	if err != nil {
		fmt.Printf("Get失败:%s\n", err.Error())
		return
	}
	if !ok {
		fmt.Println("查无结果")
		return
	}
	fmt.Printf("更新前version=%d, id=%d\n", user.Version, user.Id)
	time.Sleep(50 * time.Millisecond)

	user.UserId += 1
	if affected, err := engine.ID(user.Id).Update(user); err == nil {
		if affected > 0 {
			fmt.Printf("更新成功,version=%d,更新%d行\n", user.Version, affected)
		} else {
			fmt.Println("更新失败")
		}
	} else {
		fmt.Printf("更新失败, error=%s\n", err.Error())
	}
}
