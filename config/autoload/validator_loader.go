package autoload

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"ice_flame_gin/internal/system"
)

func ValidatorLoader() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("positive_int", system.PositiveInt)
	}
}
