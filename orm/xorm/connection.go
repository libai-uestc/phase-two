package xorm

import (
	"fmt"
	"os"
	"time"

	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

func CreateEngine(host, dbname, user, pass string, port int) *xorm.Engine {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}

	engine.SetMapper(names.GonicMapper{})

	logFile, _ := os.OpenFile("C:/Users/18101/Desktop/第二阶段/phase-two/log/xorm.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	logger := log.NewSimpleLogger(logFile)
	logger.ShowSQL(true)
	engine.SetLogger(logger)
	engine.SetLogLevel(log.LOG_INFO)

	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)
	engine.SetConnMaxLifetime(time.Hour)

	if err = engine.Ping(); err != nil {
		panic(err)
	}

	return engine

}

func CreateEngineGroup(host, dbname, user, pass string, port int) *xorm.EngineGroup {
	dsn1 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	dsn2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	dsn3 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	eg, err := xorm.NewEngineGroup("mysql", []string{
		dsn1,
		dsn2,
		dsn3,
	})
	if err != nil {
		panic(err)
	}

	if err = eg.Ping(); err != nil {
		panic(err)
	}

	eg.SetPolicy(xorm.LeastConnPolicy())

	return eg
}
