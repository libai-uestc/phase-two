package orm_test

import (
	"fmt"
	"math/rand/v2"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var (
	host   = "localhost"
	port   = 3306
	dbname = "test"
	user   = "tester"
	pass   = "123456"
)

type User struct {
	Id       int    `gorm:"primaryKey;column:id" xorm:"pk"`
	UserId   int    `gorm:"column:uid" xorm:"uid"`
	Keywords string `gorm:"column:keywords" xorm:"keywords"` // 补上该字段
	Degree   string
	CreateAt time.Time `gorm:"column:create_time;autoCreateTime" xorm:"created create_time"`
	UpdateAt time.Time `gorm:"column:update_time;autoUpdateTime" xorm:"updated update_time"`
	Gender   string
	City     string
	Province string `gorm:"-" xorm:"-"`
}

func (User) TableName() string {
	return "user"
}

func InitGorm() *gorm.DB {
	DataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	newLogger := logger.New(nil, logger.Config{LogLevel: logger.Silent})
	db, err := gorm.Open(mysql.Open(DataSourceName), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		NamingStrategy:         schema.NamingStrategy{SingularTable: true},
		Logger:                 newLogger,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func InitXorm() *xorm.Engine {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}
	engine.SetLogLevel(log.LOG_OFF)
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)
	engine.SetConnMaxLifetime(time.Hour)
	if err = engine.Ping(); err != nil {
		panic(err)
	}
	return engine
}

func TestGorm(t *testing.T) {
	db := InitGorm()
	db.Where("1=1").Delete(User{})
	const C = 100
	for i := 0; i < C; i++ {
		user := User{UserId: rand.IntN(1000000000), Degree: "本科", Gender: "男", City: "上海"}
		db = db.Create(&user)
		if db.Error != nil {
			t.Fatal(db.Error)
		}
		if db.RowsAffected != 1 {
			t.Fatalf("插入了%d行", db.RowsAffected)
		}
	}

	var users []User
	db = db.Where("degree in ?", []string{"本科", "硕士"}).Where("gender=?", "男").Where("city=?", "上海").Limit(C).Find(&users)
	if db.Error != nil {
		t.Fatal(db.Error)
	}
	if db.RowsAffected != C {
		t.Fatalf("检索到%d行", db.RowsAffected)
	}
	fmt.Println(users[0])
}

func TestXorm(t *testing.T) {
	engine := InitXorm()
	engine.Where("1=1").Delete(User{})
	const C = 100
	for i := 0; i < C; i++ {
		user := User{UserId: rand.IntN(1000000000), Degree: "本科", Gender: "男", City: "上海"}
		affected, err := engine.Insert(user)
		if err != nil {
			t.Fatal(err)
		}
		if affected != 1 {
			t.Fatalf("插入了%d行", affected)
		}
	}

	session := engine.Prepare()
	var users []User
	err := session.In("degree", []string{"本科", "硕士"}).And("gender=?", "男").Where("city=?", "上海").Limit(C).Find(&users)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != C {
		t.Fatalf("检索到%d行", len(users))
	}
	fmt.Println(user[0])
}

func BenchmarkCreateGorm(b *testing.B) {
	db := InitGorm()
	db.Where("1=1").Delete(User{})

	for b.Loop() {
		user := User{UserId: rand.IntN(1000000000), Degree: "本科", Gender: "男", City: "上海"}
		db.Create(&user)
	}
}

func BenchmarkCreateXorm(b *testing.B) {
	engine := InitXorm()
	engine.Where("1=1").Delete(User{})

	for b.Loop() {
		user := User{UserId: rand.IntN(1000000000), Degree: "本科", Gender: "男", City: "上海"}
		engine.Insert(user)
	}
}

func BenchmarkReadGorm(b *testing.B) {
	db := InitGorm()
	const C = 100

	for b.Loop() {
		var users []User
		db.Where("degree in ?", []string{"本科", "硕士"}).Where("gender=?", "男").Where("city=?", "上海").Limit(C).Find(&users)
	}
}

func BenchmarkReadXorm(b *testing.B) {
	engine := InitXorm()
	const C = 100
	session := engine.Prepare()

	for b.Loop() {
		var users []User
		session.In("degree", []string{"本科", "硕士"}).Where("gender=?", "男").Where("city=?", "上海").Limit(C).Find(&users)
	}
}
