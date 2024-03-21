package autoload

import (
	"github.com/go-playground/validator/v10"
	"ice_flame_gin/internal/system"
	"log"
)

func ValidatorLoader() {
	validate := validator.New()
	err := validate.RegisterValidation("positive_int", system.PositiveInt)
	if err != nil {
		log.Fatalf("注册自定义验证函数失败：%v", err)
	}
}
