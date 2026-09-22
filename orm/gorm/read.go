package gorm

import (
	"errors"
	"fmt"
	"log/slog"

	"gorm.io/gorm"
	"gorm.io/hints"
)

func Read(db *gorm.DB) {
	user := User{City: "HongKong", Id: 3}
	tx := db.Select("uid,city,gender,keywords").Where("uid>100 and degree='大专'").Where("city in ?", []string{"北京", "上海"}).Where("degree like ?", "%科").Or("gender=?", "女").Order("id desc, uid").Order("city").Offset(3).Limit(1).First(&user)

	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			slog.Error("读DB失败", "error", tx.Error)
		} else {
			slog.Info("查无结果")
		}
	} else {
		if tx.RowsAffected > 0 {
			fmt.Printf("read结果：%+v\n", user)
		} else {
			slog.Info("查无结果", "user", user)
		}
	}

	var user2 *User
	tx = db.Find(user2)

	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			slog.Error("读DB失败", "error", tx.Error)
		} else {
			slog.Info("查无结果")
		}
	}

	var user3 *User = new(User)
	tx = db.Find(user3)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			slog.Error("读DB失败", "error", tx.Error)
		} else {
			slog.Info("查无结果")
		}
	} else {
		if tx.RowsAffected > 0 {
			fmt.Printf("read结果：%+v\n", user3)
		} else {
			slog.Info("查无结果", "user", user3)
		}
	}

	var users []User
	tx = db.Limit(3).Find(&users)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			slog.Error("读DB失败", "error", tx.Error)
		} else {
			slog.Info("查无结果")
		}
	} else {
		if tx.RowsAffected > 0 {
			fmt.Println("多个read结果")
			for _, u := range users {
				fmt.Printf("%+v\n", u)
			}
		} else {
			slog.Info("查无结果")
		}
	}

	user4 := User{Id: 23212}
	tx = db.Find(&user4)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			slog.Error("读DB失败", "error", tx.Error)
		} else {
			slog.Info("查无结果")
		}
	} else {
		if tx.RowsAffected > 0 {
			fmt.Printf("read结果：%+v\n", user4)
		} else {
			slog.Info("查无结果")
		}
	}

	db.Where("uid>0").Clauses(hints.UseIndex("id", "idx_uid")).Find(&users)
	db.Where("uid>0").Clauses(hints.ForceIndex("idx_uid")).Find(&users)
}

func ReadWithStatistics(db *gorm.DB) {
	type Result struct {
		City string
		Mid  float64
	}

	var results []Result
	db.Model(User{}).Select("city,avg(id) as mid").Group("city").Having("mid>0").Find(&results)
	fmt.Println("group by having查询结果：")
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}

	db.Table("user").Distinct("city").Find(&results)
	fmt.Println("distinct查询结果：")
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}

	var count int64
	db.Table("user").Where("city=?", "北京").Count(&count)
	fmt.Printf("count=%d\n", count)
}
