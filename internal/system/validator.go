package system

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
	"strings"
)

// GetValidationErrors
//
// @Title GetValidationErrors
// @Description: 获取验证错误信息
// @Author liuxingyu
// @Date 2024-02-09 01:59:07
// @param err
// @param form
// @return map[string]string
func GetValidationErrors(err error, form interface{}) map[string]string {
	errorMap := make(map[string]string)

	// 判断是否为验证错误
	if v, ok := err.(validator.ValidationErrors); ok {
		for _, e := range v {
			fieldName := e.Field()
			tagName := e.Tag()

			// 根据字段和标签获取自定义错误消息
			msg := GetCustomErrorMsg(fieldName, tagName, form)

			// 将错误信息添加到映射中
			errorMap[fieldName] = msg
		}
	} else {
		// 非验证错误，直接使用默认错误信息
		errorMap["non-field-error"] = err.Error()
	}

	return errorMap
}

// GetValidationErrorMsg
//
// @Title GetValidationErrorMsg
// @Description: 获取验证错误信息
// @Author liuxingyu
// @Date 2024-01-31 23:51:02
// @param err
// @param form
// @return string
func GetValidationErrorMsg(err error, form interface{}) string {
	errMsg := ""

	// 判断是否为验证错误
	if v, ok := err.(validator.ValidationErrors); ok {
		for _, e := range v {
			fieldName := e.Field()
			tagName := e.Tag()

			// 根据字段和标签获取自定义错误消息
			msg := GetCustomErrorMsg(fieldName, tagName, form)

			// 将错误信息拼接到一起
			errMsg += msg + ", "
		}

		// 去除最后一个逗号和空格
		errMsg = strings.TrimSuffix(errMsg, ", ")
	} else {
		// 非验证错误，直接使用默认错误信息
		errMsg = err.Error()
	}

	return errMsg
}

// GetCustomErrorMsg
//
// @Title GetCustomErrorMsg
// @Description: 根据字段和标签获取自定义错误消息
// @Author liuxingyu
// @Date 2024-01-31 23:50:45
// @param fieldName
// @param tagName
// @param form
// @return string
func GetCustomErrorMsg(fieldName string, tagName string, form interface{}) string {
	// 使用反射获取结构体字段上的标签信息
	t := reflect.TypeOf(form)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	f, _ := t.FieldByName(fieldName)
	tag := f.Tag.Get("binding")
	// 解析标签中的验证条件
	tagConditions := strings.Split(tag, ",")

	tagMsg := f.Tag.Get("msg")
	// 解析标签中的错误消息
	tagMsgs := strings.Split(tagMsg, "|")

	// 根据字段名和具体验证条件进行匹配
	for i, condition := range tagConditions {
		// 检查是否包含字段名和验证条件
		if strings.Contains(condition, tagName) {
			// 返回匹配的错误消息
			return tagMsgs[i]
		}
	}

	// 默认返回标签中的错误消息
	if len(tagMsgs) > 0 {
		return tagMsgs[0]
	}

	// 默认返回默认错误消息
	return tagName
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
	field := fl.Field()
	fmt.Println(field)
	switch field.Kind() {
	case reflect.String:
		val, err := strconv.Atoi(field.String())
		if err != nil {
			return false
		}
		return val > 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() > 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return field.Uint() > 0
	default:
		return false
	}
}
