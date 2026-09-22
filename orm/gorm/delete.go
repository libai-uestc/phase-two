package gorm

import (
	"fmt"

	"gorm.io/gorm"
)

func Delete(db *gorm.DB) {
	tx := db.Where("degree=?", "本科").Delete(User{})
	fmt.Printf("删除%d行\n", tx.RowsAffected)
	var user User = User{Id: 10}
	db.Delete(user)

	db.Delete(User{}, 1)
	db.Delete(User{}, []int{1, 2, 3})
}
