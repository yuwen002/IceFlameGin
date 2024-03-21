package autoload

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"ice_flame_gin/internal/system"
	"log"
)

func ValidatorLoader() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("positive_int", system.PositiveInt)
		if err != nil {
			log.Fatalf("注册自定义验证函数失败：%v", err)
		}
	}
}
