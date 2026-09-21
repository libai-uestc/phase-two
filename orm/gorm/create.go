package gorm

import (
	"fmt"
	"log/slog"
	"math/rand/v2"

	"gorm.io/gorm"
)

func Create(db *gorm.DB) {
	user := User{UserId: rand.IntN(100000), Degree: "本科", Gender: "男", City: "上海", Keywords: []string{"编程", "golang"}}
	result := db.Create(&user)
	if result.Error != nil {
		slog.Error("插入记录失败", "error", result.Error)
	}
	fmt.Printf("record id is %d\n", user.Id)
	fmt.Printf("影响%d行\n", result.RowsAffected)
}
