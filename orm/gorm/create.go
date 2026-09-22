package gorm

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"time"

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

	tx := db.Session(&gorm.Session{SkipHooks: true})

	user1 := user
	user1.Id = 0
	user1.UserId = rand.IntN(100000)
	user2 := user
	user2.Id = 0
	user2.UserId = rand.IntN(100000)
	users := []*User{&user1, &user2}
	result = tx.Create(users)
	fmt.Printf("影响%d行\n", result.RowsAffected)

	batchSize := 1
	user3 := user
	user3.Id = 0
	user3.UserId = rand.IntN(100000)
	user4 := user3
	user4.Id = 0
	db.CreateInBatches([]*User{&user3, &user4}, batchSize)
}

func CreateByMap(db *gorm.DB) error {
	err := db.Model(User{}).Create(map[string]any{
		"uid": rand.IntN(100000), "degree": "本科",
		"gender": "男", "city": "上海",
		"create_time": time.Now(),
		"keywords":    "[]",
	}).Error
	if err != nil {
		return err
	}

	err = db.Model(User{}).Create([]map[string]any{
		{"uid": rand.IntN(100000), "degree": "本科", "gender": "男", "city": "北京", "create_time": time.Now(), "keywords": "[]"},
		{"uid": rand.IntN(100000), "degree": "本科", "gender": "男", "city": "深圳", "create_time": time.Now(), "keywords": "[]"},
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeSave(db *gorm.DB) (err error) {
	db.Logger.Info(context.Background(), "exec hook BeforeSave")
	return nil
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	db.Logger.Info(context.Background(), "exec hook BeforeCreate")
	return nil
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	db.Logger.Info(context.Background(), "exec hook AfterCreate")
	return nil
}

func (u *User) AfterSave(db *gorm.DB) (err error) {
	db.Logger.Info(context.Background(), "exec hook AfterSave")
	return nil
}
