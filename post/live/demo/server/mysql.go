package main

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB // 连接池
	once sync.Once
)

var (
	user   = "tester"
	pass   = "123456"
	host   = "localhost"
	port   = 3306
	dbName = "test"
)

// func init(){

// }

func GetDB() *gorm.DB {
	// if db == nil {
	// 	DataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbName)
	// 	var err error
	// 	db, err = gorm.Open(mysql.Open(DataSourceName),
	// 		&gorm.Config{PrepareStmt: true},
	// 	)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	once.Do(func() {
		DataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbName)
		var err error
		db, err = gorm.Open(mysql.Open(DataSourceName),
			&gorm.Config{PrepareStmt: true},
		)
		if err != nil {
			panic(err)
		}
	})
	return db
}

func CloseDB() {
	if db != nil {
		sqlDB, _ := db.DB()
		sqlDB.Close() //关闭连接池
	}
}

// users
type User struct {
	// Name     string `gorm:"column:username"`
	Name     string
	PassWord string `gorm:"column:password"` // pass_word
}

func (User) TableName() string {
	return "login"
}

func GetPasswordByName(name string) (string, error) {
	var user User
	// GetDB().Select("password","name").Where("name=?",name).First(&user)
	// GetDB().Select("password").Where("name=?",name).First(&user) // ?可以防止sql注入攻击
	// GetDB().Select("password").Where("name=?",name).Find(&user) // 切片用find
	// 如果只有一个结果,可以用Take或First
	err := GetDB().Select("password").Where("name=?", name).Take(&user).Error
	if err != nil {
		err2 := fmt.Errorf("GetPasswordByName %w", err)
		return "", err2
	}
	return user.PassWord, err
}
