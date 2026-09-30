package handler

import (
	database "libai/go/phase-two/post/database/gorm"
	"libai/go/phase-two/post/handler/model"
	"libai/go/phase-two/post/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 从业务上来说，跟用户相关的：用户注册、用户登录、用户修改密码

// 用户注册，要从前端拿到用户输入的用户名和密码，调用database，把数据进行入库
// 本质上是获得http请求参数，通过参数绑定来完成，参数绑定可以同时完成参数的校验
func RegistUser(ctx *gin.Context) {
	// name := ctx.PostForm("name")
	// pass := ctx.PostForm("pass")
	var user model.User
	err := ctx.ShouldBind(&user)
	// 但如果要做一些数据验证，如密码长度是否在某个区间内，用户名的长度必须要大于0或大于4
	if err != nil {
		ctx.String(http.StatusBadRequest, util.BindErrMsg(err))
		return
	}
	_, err = database.RegistUser(user.Name, user.PassWord)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	// ctx.Status(200)

}

// handler里的model，不是和数据库进行映射了，主要是请求参数封装成一个结构体

func Login(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, util.BindErrMsg(err))
		return
	}

	user2 := database.GetUserByName(user.Name)
	if user2 == nil {
		ctx.String(http.StatusBadRequest, "用户名不存在")
		return
	}

	if user2.PassWord != user.PassWord {
		ctx.String(http.StatusBadRequest, "密码错误")
		return
	}

	// 登录成功，返回cookie
	ctx.SetCookie("uid", strconv.Itoa(user2.Id), 86400, "/", "localhost", false, true)
}

func Logout(ctx *gin.Context) {
	ctx.SetCookie("uid", "", -1, "/", "localhost", false, true)
}

func UpdatePassword(ctx *gin.Context) {
	var req model.ModifyPassRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.String(http.StatusBadRequest, util.BindErrMsg(err))
		return
	}

	uid := GetUidFromCookie(ctx)
	if uid <= 0 {
		ctx.String(http.StatusForbidden, "请先登录")
		return
	}

	err = database.UpdatePassword(uid, req.NewPass, req.OldPass)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

}

func GetUidFromCookie(ctx *gin.Context) int {
	for _, cookie := range ctx.Request.Cookies() {
		if cookie.Name == "uid" {
			uid, err := strconv.Atoi(cookie.Value)
			if err == nil {
				return uid
			}
		}
	}
	return 0
}

// func Login(ctx *gin.Context) {
// 	var user model.User
// 	err := ctx.ShouldBind(&user)
// 	if err != nil {
// 		ctx.String(http.StatusBadRequest, util.BindErrMsg(err))
// 		return
// 	}

// 	user2 := database.GetUserByName(user.Name)
// 	if user2 == nil {
// 		ctx.String(http.StatusBadRequest, "用户名不存在")
// 		return
// 	}
// 	if user2.PassWord != user.PassWord {
// 		ctx.String(http.StatusBadRequest, "密码错误")
// 		return
// 	}

// 	slog.Info("登录成功", "uid", user2.Id)

// }
