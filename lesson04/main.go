package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// user info
type User struct {
	ID     int64  `json:"id" validate:"gt=0"`
	Name   string `json:"name" validate:"required"`
	Gender string `json:"gender" validate:"required,oneof=man woman"`
	Age    uint8  `json:"age" validate:"required,gte=0,lte=130"`
	Email  string `json:"email" validate:"required,email"`
}

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
)

func main() {
	validate = validator.New()

	// 验证变量
	email := ""
	err := validate.Var(email, "required,email")
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		fmt.Println(validationErrors)
		// output: Key: '' Error:Field validation for '' failed on the 'email' tag
		// output: Key: '' Error:Field validation for '' failed on the 'required' tag
		return
	}

	// 验证结构体
	// 注册一个函数，获取结构字段的备用名称。
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return "j"
		}
		return name
	})

	// 翻译为中文
	zh := zh.New()
	uni = ut.New(zh)
	trans, _ := uni.GetTranslator("zh")
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)
	user := &User{
		ID:     1,
		Name:   "lucy",
		Gender: "boy",
		Age:    180,
		Email:  "gopher@88.com",
	}
	err = validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// output: Key: 'User.Age' Error:Field validation for 'Age' failed on the 'lte' tag
		// fmt.Println(validationErrors)
		fmt.Println(validationErrors.Translate(trans))
		return
	}
}
