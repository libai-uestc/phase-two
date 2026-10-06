package xorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Login struct {
	Username string
	Password string
}

// func XormQuickStart() {
// 	host := "localhost"
// 	port := 3306
// 	dbname := "test"
// 	user := "tester"
// 	pass := "123456"

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbname)
// 	engine, err := xorm.NewEngine("mysql", dsn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer engine.Close()
// 	if err = engine.Ping(); err != nil {
// 		panic(err)
// 	}
// 	instance1 := Login{Username: "libai", Password: "123456"}
// 	engine.Insert(&instance1)

//		var instance2 Login
//		engine.Get(&instance2)
//		fmt.Printf("%#v\n", instance2)
//	}
func XormQuickStart() {
	host := "localhost"
	port := 3306
	dbname := "test"
	user := "tester"
	pass := "123456"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbname)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer engine.Close()

	if err = engine.Ping(); err != nil {
		panic(err)
	}

	// 【新增 1】：自动同步表结构（如果 login 表不存在，会自动在数据库创建）
	if err := engine.Sync2(new(Login)); err != nil {
		panic(fmt.Errorf("建表失败: %v", err))
	}

	// 【修改 2】：接收 Insert 的错误并检查
	instance1 := Login{Username: "libai", Password: "123456"}
	if _, err := engine.Insert(&instance1); err != nil {
		panic(fmt.Errorf("插入失败: %v", err))
	}

	// 【修改 3】：接收 Get 的返回值并检查
	var instance2 Login
	// 建议：给 Get 加一个查询条件，否则如果表里有多条数据，它不知道查哪条
	has, err := engine.Where("username = ?", "libai").Get(&instance2)
	if err != nil {
		panic(fmt.Errorf("查询失败: %v", err))
	}

	if has {
		fmt.Printf("查到了数据: %#v\n", instance2)
	} else {
		fmt.Println("数据库中没有这条记录")
	}
}
