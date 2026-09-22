package gorm

import (
	"fmt"
	"math/rand/v2"

	"gorm.io/gorm"
)

// func Save(db *gorm.DB) {
// 	db = db.Debug()
// 	user := User{UserId: rand.IntN(100000), Degree: "本科", Gender: "男", City: "上海"}
// 	db.Save(&user)

// 	var user2 User
// 	db.Last(&user2)
// 	user2.Degree = "硕士"
// 	db.Save(&user2)
// }

func Save(db *gorm.DB) {
	// 1. 检查是不是一上来连接就是坏的
	if db.Error != nil {
		fmt.Printf("🚨 数据库连接本身就有错误，请检查 CreateConnection: %v\n", db.Error)
		return
	}

	db = db.Debug() // 开启日志

	// 2. 尝试插入数据
	user := User{UserId: rand.IntN(100000), Degree: "本科", Gender: "男", City: "上海", Keywords: []string{}}
	err := db.Save(&user).Error
	if err != nil {
		fmt.Printf("❌ 插入数据失败: %v\n", err)
		return
	}
	fmt.Println("✅ 第一步：插入数据成功！")

	// 3. 尝试查询最后一条数据
	var user2 User
	err = db.Last(&user2).Error
	if err != nil {
		fmt.Printf("❌ 查询最后一条数据失败: %v\n", err)
		return
	}
	fmt.Println("✅ 第二步：查询最后一条成功，查到的主键 ID 是:", user2.Id)

	// 4. 尝试更新数据
	user2.Degree = "硕士"
	err = db.Save(&user2).Error
	if err != nil {
		fmt.Printf("❌ 更新数据失败: %v\n", err)
		return
	}
	fmt.Println("✅ 第三步：更新数据成功！")
}

func Update(db *gorm.DB) {
	tx := db.Model(&User{}).
		Where("city=?", "北京").Updates(
		map[string]any{"degree": "硕士", "gender": "男"},
	)
	fmt.Printf("更新了%d行\n", tx.RowsAffected)
	db.Model(&User{}).Where("city=?", "北京").Updates(User{Degree: "本科", Gender: "男", Id: 1})
	fmt.Printf("更新了%d行\n", tx.RowsAffected)
}
