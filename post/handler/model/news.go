package model

import "time"

// type News struct {
// 	Id           int `gorm:"primaryKey" xorm:"pk antoincr"`
// 	UserId       int
// 	UserName     string `gorm:"-" xorm:"-"`
// 	Title        string
// 	Content      string     `gorm:"column:article"`
// 	PostTime     *time.Time `gorm:"column:create_time" xorm:"create_time created"`
// 	DeleteTime   *time.Time `xorm:"deleted"`
// 	ViewPostTime string     `gorm:"-" xorm:"-"`
// }
type News struct {
	Id           int `gorm:"primaryKey" xorm:"pk autoincr"` // 修复了 autoincr 拼写
	UserId       int
	UserName     string     `gorm:"-" xorm:"-"`
	Title        string     `form:"title"`                         // 👈 必须加这个
	Content      string     `gorm:"column:article" form:"content"` // 👈 必须加这个
	PostTime     *time.Time `gorm:"column:create_time" xorm:"create_time created"`
	DeleteTime   *time.Time `xorm:"deleted"`
	ViewPostTime string     `gorm:"-" xorm:"-"`
}
