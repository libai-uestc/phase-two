package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// orm:结构体->表
// mysql -u tester -p
// 123456
// use test;
// show tables;
// desc login;

type Login struct {
	Username string
	Password string
}

func (Login) TableName() string {
	return "login"
}

func GormQuickStart() {
	host := "localhost"
	port := 3306
	dbname := "test"
	user := "tester"
	pass := "123456"
	// data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		panic(err)
	}

	// 写入数据
	instance1 := Login{Username: "libai", Password: "123456"}
	db.Create(&instance1)

	// 读取
	var instance2 Login
	db.Find(&instance2)
	fmt.Printf("%#v\n", instance2)
	// select * from login;
}
