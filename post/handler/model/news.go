package model

import "time"

type News struct {
	Id           int `gorm:"primaryKey" xorm:"pk antoincr"`
	UserId       int
	UserName     string `gorm:"-" xorm:"-"`
	Title        string
	Content      string     `gorm:"column:article"`
	PostTime     *time.Time `gorm:"column:create_time" xorm:"create_time created"`
	DeleteTime   *time.Time `xorm:"deleted"`
	ViewPostTime string     `gorm:"-" xorm:"-"`
}
