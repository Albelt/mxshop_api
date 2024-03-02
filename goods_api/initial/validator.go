package initial

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"mxshop_api/goods_api/model"
)

// 注册自定义的验证器
func InitValidator() {
	var (
		err error
	)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// mobile验证器
		if err = v.RegisterValidation("mobile", model.ValidateMobile); err != nil {
			panic(err)
		}
	}
}
