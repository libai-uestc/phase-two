package gorm

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func HandleError(db *gorm.DB) {
	var user User
	db.First(&user)

	tx := db.Create(&user)
	if tx.Error != nil {
		if mysqlErr, ok := tx.Error.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1:
			case 2:
			default:
				fmt.Println("mysql error", "code", mysqlErr.Number, "msg", mysqlErr.Message)
			}
		} else {
			fmt.Printf("err %#v\n", tx.Error)
		}
	}
	if tx.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(tx.Error, &mysqlErr) {
			switch mysqlErr.Number {
			case 1:
			case 2:
			default:
				fmt.Println("mysql error", "code", mysqlErr.Number, "msg", mysqlErr.Message)
			}
		} else {
			fmt.Printf("err %#v\n", tx.Error)
		}
	}
}
