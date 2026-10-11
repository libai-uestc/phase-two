package xorm

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func HandleError(engine *xorm.Engine) {
	var user User
	engine.Get(&user)

	_, err := engine.Insert(user)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1:
				// ...
			case 2:
				// ....
			default:
				fmt.Println("mysql error", "code", mysqlErr.Number, "msg", mysqlErr.Message)
			}
		} else {
			fmt.Printf("err %#v\n", err)
		}
	}

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1:
				// ...
			case 2:
				// ....
			default:
				fmt.Println("mysql error", "code", mysqlErr.Number, "msg", mysqlErr.Message)
			}
		}
	}
}
