package xorm

import (
	"fmt"
	"log/slog"

	"xorm.io/xorm"
)

func Delete(engine *xorm.Engine) {
	affected, err := engine.ID(10).Delete(User{})
	if err != nil {
		slog.Error("删除记录失败", "error", err)
	}
	fmt.Printf("删除%d行\n", affected)

	affected, err = engine.Where("degree=?", "本科").Delete(User{})
	if err != nil {
		slog.Error("删除记录失败", "error", err)
	}
	fmt.Printf("删除%d行\n", affected)
}
