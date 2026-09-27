package database

import (
	"errors"
	"fmt"
	"libai/go/phase-two/post/database/model"
	"log/slog"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func RegistUser(name, password string) (int, error) {
	user := model.User{
		Name:     name,
		PassWord: password,
	}
	err := PostDB.Create(&user).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 {
				return 0, fmt.Errorf("用户名[%s]已存在", name)
			}
		}
		slog.Error("用户注册失败", "name", name, "error", err)
		return 0, errors.New("用户注册失败，请稍后重试")
	}
	return user.Id, nil
}

func LogOffUser(uid int) error {
	user := model.User{Id: uid}
	tx := PostDB.Delete(user)
	if tx.Error != nil {
		slog.Error("注销用户失败", "uid", uid, "error", tx.Error)
		return errors.New("用户注销失败，请稍后重试")
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("用户注销失败，uid %d不存在", uid)
	}
	return nil
}

func GetUserById(uid int) *model.User {
	user := &model.User{Id: uid}
	tx := PostDB.Select("*").First(user)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			slog.Error("GetUserById failed", "uid", uid, "error", tx.Error)
		}
		return nil
	}
	return user
}

func GetUserByName(name string) *model.User {
	user := &model.User{}
	tx := PostDB.Select("*").Where("name=?", name).First(user)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			slog.Error("GetUserByName failed", "name", name, "error", tx.Error)
		}
		return nil
	}
	return user
}

func UpdateUserName(uid int, name string) error {
	tx := PostDB.Model(&model.User{}).Where("id=?", uid).Update("name", name)
	if tx.Error != nil {
		slog.Error("UpdateUserName failed", "uid", uid, "new name", name, "error", tx.Error)
		return errors.New("用户名修改失败，请稍后重试")
	} else {
		if tx.RowsAffected <= 0 {
			return fmt.Errorf("用户id[%d]不存在", uid)
		} else {
			return nil
		}
	}
}

func UpdatePassword(uid int, newPass, oldPass string) error {
	tx := PostDB.Model(&model.User{}).Where("id=? and password=?", uid, oldPass).Update("password", newPass)
	if tx.Error != nil {
		slog.Error("UpdatePassword failed", "uid", uid, "error", tx.Error)
		return errors.New("密码修改失败，请稍后重试")
	} else {
		if tx.RowsAffected <= 0 {
			return errors.New("旧密码不对")
		} else {
			return nil
		}
	}
}
