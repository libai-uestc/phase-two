package xorm

import (
	"fmt"

	"xorm.io/xorm"
)

func RawSelect(engine *xorm.Engine) {
	results, err := engine.QueryInterface("select id,uid,city from user where id>? and uid>? limit 3", 2, 4)
	if err == nil {
		for _, result := range results {
			fmt.Printf("%+v\n", result)
		}
	}
	fmt.Println()
}

func RawExec(engine *xorm.Engine) {
	result, err := engine.Exec("update user set degree=? where id=?", "大专", 24)
	if err == nil {
		affected, _ := result.RowsAffected()
		fmt.Printf("更新了%d行\n", affected)
	}
}
