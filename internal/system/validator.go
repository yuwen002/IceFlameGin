package system

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

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
	tag := f.Tag.Get("msg")

	// 解析标签中的错误消息
	tagMsgs := strings.Split(tag, "|")
	for _, msg := range tagMsgs {
		msgParts := strings.SplitN(msg, ":", 2)
		if len(msgParts) == 2 && msgParts[0] == tagName {
			return msgParts[1]
		}
	}

	// 默认返回标签中的错误消息
	if len(tagMsgs) > 0 {
		return tagMsgs[0]
	}

	// 默认返回默认错误消息
	return tagName
}
