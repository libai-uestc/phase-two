package xorm

import "time"

type User struct {
	Id         int `xorm:"pk autoincr"`
	UserId     int `xorm:"uid"`
	Degree     string
	Keywords   []string  `xorm:"json"`
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
	DeleteTime time.Time `xorm:"deleted"`
	Gender     string
	City       string
	Version    int    `xorm:"version"`
	Province   string `xorm:"-"`
}
