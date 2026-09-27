package model

type User struct {
	Id       int `gorm:"primarKey"`
	Name     string
	PassWord string `gorm:"column:password"`
}
