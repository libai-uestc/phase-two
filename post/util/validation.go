package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func BindErrMsg(err error) string {
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
		return fmt.Sprintf("invalid error type: %#v", err)
	}
}
