package xorm

import (
	"fmt"
	"math/rand/v2"

	"xorm.io/xorm"
)

func Transaction(engine *xorm.Engine) error {
	session := engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	user := User{UserId: rand.IntN(100000), Degree: "本科", Gender: "男", City: "上海"}
	fmt.Printf("uid=%d\n", user.UserId)
	if _, err := session.Insert(&user); err != nil {
		session.Rollback()
		fmt.Println("第一次Insert回滚")
		return err
	}
	// user.Id = 0
	// if _, err := session.Insert(&user); err != nil { //第二次会失败，因为uid重复了
	// 	session.Rollback() //手动回滚
	// 	fmt.Println("第二次Insert回滚")
	// 	return err
	// }

	//提交事务
	// session.Commit()

	user = User{UserId: rand.IntN(100000), Degree: "本科", Gender: "男", City: "上海"}
	fmt.Printf("uid=%d\n", user.UserId)
	if _, err := session.Insert(&user); err != nil {
		session.Rollback()
		fmt.Println("第三次Insert回滚")
		return err
	}

	return session.Commit()
}
