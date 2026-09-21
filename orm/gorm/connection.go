package gorm

import (
	"fmt"
	"log"
	"os"
	// "runtime/trace"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func CreateConnection(host, dbname, user, pass string, port int) *gorm.DB {
	DataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	logFile, _ := os.OpenFile("C:/Users/18101/Desktop/第二阶段/phase-two/log/gorm.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             500 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(mysql.Open(DataSourceName), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger:                   newLogger,
		DryRun:                   false,
		DisableNestedTransaction: true,
		DisableAutomaticPing:     false,
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
