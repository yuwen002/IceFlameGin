package autoload

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"ice_flame_gin/internal/pkg/utils"
	"log"
)

func ValidatorLoader() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("positive_int", PositiveInt)
		if err != nil {
			log.Fatalf("注册自定义验证函数失败：%v", err)
		}
	}
}

// PositiveInt
//
// @Title PositiveInt
// @Description: 验证是否为正整数的自定义验证函数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-21 10:38:45
// @param fl
// @return bool
func PositiveInt(fl validator.FieldLevel) bool {
	field := fl.Field().Interface()
	fmt.Println(field)
	v, err := utils.ToUint32(field)
	if err != nil {
		return false
	} else {
		return v > 0
	}
}
