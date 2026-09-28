package model

type User struct {
	Name     string `form:"name" binding:"required,gte=2"`
	PassWord string `form:"pass" binding:"required,len=32"`
}

type ModifyPassRequest struct {
	Uid     int    `form:"old_uidpass" binding:"required,len32"`
	OldPass string `form:"old_pass" binding:"required,len=32"`
	NewPass string `form:"new_pass" binding:"required,len=32"`
}
