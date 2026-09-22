package gorm

import (
	"fmt"
	"math/rand/v2"

	"gorm.io/gorm"
)

func Transaction(db *gorm.DB) error {
	db = db.Session(&gorm.Session{SkipDefaultTransaction: true})

	tx := db.Begin()

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	user := User{UserId: rand.IntN(100000), Degree: "本科", Gender: "男", City: "上海"}
	fmt.Printf("uid=%d\n", user.UserId)
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		fmt.Println("第一次Create回滚")
		return err
	}
	user.Id = 0
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		fmt.Println("第二次Create回滚")
		return err
	}
	return tx.Commit().Error
}
