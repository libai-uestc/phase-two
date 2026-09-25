package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `form:"name" binding:"required"`
	Score int    `form:"score" binding:"gt=0,required"`

	Enrollment time.Time `form:"enrollment" binding:"required,before_today" time_format:"2006-01-02" time_utc:"8"`
	Graduation time.Time `form:"graduation" binding:"required,gtfield=Enrollment" time_format:"2006-01-02" time_utc:"8"`
}

var beforeToday validator.Func = func(f1 validator.FieldLevel) bool {
	if date, ok := f1.Field().Interface().(time.Time); ok {
		today := time.Now()
		if date.Before(today) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func processErr(err error) string {
	if err == nil {
		return ""
	}

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		msgs := make([]string, 0, 3)
		for _, validationErr := range validationErrs {
			msgs = append(msgs, fmt.Sprintf("字段 [%s] 不满足条件[%s]", validationErr.Field(), validationErr.Tag()))
		}
		return strings.Join(msgs, ";")
	} else {
		return "invalid error"
	}
}

func main() {
	engine := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("before_today", beforeToday)
	}

	engine.GET("/", func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBind(&user); err != nil {
			msg := processErr(err)
			ctx.String(http.StatusBadRequest, "参数绑定失败："+msg)
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	})

	engine.Run("127.0.0.1:5678")
}
