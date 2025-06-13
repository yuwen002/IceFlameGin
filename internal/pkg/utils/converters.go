package utils

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

// ToString
//
// @Title ToString
// @Description: 将任意类型转换为string
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 15:01:13
// @param value
// @return string
// @return error
func ToString(value interface{}) (string, error) {
	switch v := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return fmt.Sprintf("%v", v), nil
	case string:
		return v, nil
	case bool:
		return fmt.Sprintf("%v", v), nil
	case []byte:
		return string(v), nil
	case time.Time:
		return v.Format(time.RFC3339), nil
	default:
		// 检查是否是指向 time.Time 类型的指针
		if reflect.TypeOf(value).Kind() == reflect.Ptr {
			val := reflect.ValueOf(value).Elem()
			if val.Kind() == reflect.Struct && val.Type().String() == "time.Time" {
				t := val.Interface().(time.Time)
				return t.Format(time.RFC3339), nil
			}
		}
		return "", errors.New("不支持的类型")
	}
}

// ToInt
//
// @Title ToInt
// @Description: 将任意类型转换为整数类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 15:07:02
// @param value
// @return int
// @return error
func ToInt(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		return int(v), nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		if v == "" {
			return 0, errors.New("空字符串无法转换为int")
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return i, nil
	default:
		return 0, errors.New("不支持的类型")
	}
}

// ToInt8
//
// @Title ToInt8
// @Description: 将给定的值转换为 int8 类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 15:30:18
// @param value
// @return int8
// @return error
func ToInt8(value interface{}) (int8, error) {
	switch v := value.(type) {
	case int8:
		return v, nil
	case int:
		if v > 127 || v < -128 {
			return 0, errors.New("超出int8范围")
		}
		return int8(v), nil

	case int16:
		if v > 127 || v < -128 {
			return 0, errors.New("超出int8范围")
		}
		return int8(v), nil
	case int32:
		if v > 127 || v < -128 {
			return 0, errors.New("超出int8范围")
		}
		return int8(v), nil
	case int64:
		if v > 127 || v < -128 {
			return 0, errors.New("超出int8范围")
		}
		return int8(v), nil
	case uint:
		if v > 127 {
			return 0, errors.New("超出int8范围")
		}
		return int8(v), nil
	case uint8:
		return int8(v), nil
	case uint16:
		if v > 127 {
			return 0, errors.New("超出int8范围")
		}
		return int8(v), nil
	case uint32:
		if v > 127 {
			return 0, errors.New("超出int8范围")
		}
		return int8(v), nil
	case uint64:
		if v > 127 {
			return 0, errors.New("超出int8范围")
		}
		return int8(v), nil
	case float32:
		if v > 127 || v < -128 {
			return 0, errors.New("超出int8范围")
		}
		return int8(v), nil
	case float64:
		if v > 127 || v < -128 {
			return 0, errors.New("超出int8范围")
		}
		return int8(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		if v == "" {
			return 0, errors.New("空字符串无法转换为int8")
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		if i > 127 || i < -128 {
			return 0, errors.New("超出int8范围")
		}
		return int8(i), nil
	default:
		return 0, errors.New("不支持的类型")
	}
}

// ToInt16
//
// @Title ToInt16
// @Description: 将给定的值转换为 int16 类型。
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 15:34:00
// @param value
// @return int16
// @return error
func ToInt16(value interface{}) (int16, error) {
	switch v := value.(type) {
	case int:
		if v > 32767 || v < -32768 {
			return 0, errors.New("超出 int16 范围")
		}
		return int16(v), nil
	case int8:
		return int16(v), nil
	case int16:
		return v, nil
	case int32:
		if v > 32767 || v < -32768 {
			return 0, errors.New("超出 int16 范围")
		}
		return int16(v), nil
	case int64:
		if v > 32767 || v < -32768 {
			return 0, errors.New("超出 int16 范围")
		}
		return int16(v), nil
	case uint:
		if v > 32767 {
			return 0, errors.New("超出 int16 范围")
		}
		return int16(v), nil
	case uint8:
		return int16(v), nil
	case uint16:
		return int16(v), nil
	case uint32:
		if v > 32767 {
			return 0, errors.New("超出 int16 范围")
		}
		return int16(v), nil
	case uint64:
		if v > 32767 {
			return 0, errors.New("超出 int16 范围")
		}
		return int16(v), nil
	case float32:
		if v > 32767 || v < -32768 {
			return 0, errors.New("超出 int16 范围")
		}
		return int16(v), nil
	case float64:
		if v > 32767 || v < -32768 {
			return 0, errors.New("超出 int16 范围")
		}
		return int16(v), nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		i, err := strconv.ParseInt(v, 10, 16)
		if err != nil {
			return 0, err
		}
		return int16(i), nil
	default:
		return 0, errors.New("不支持的类型")
	}
}

// ToInt32
//
// @Title ToInt32
// @Description: 将给定的值转换为 int32 类型。
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 15:36:33
// @param value
// @return int32
// @return error
func ToInt32(value interface{}) (int32, error) {
	switch v := value.(type) {
	case int:
		if v > 2147483647 || v < -2147483648 {
			return 0, errors.New("超出 int32 范围")
		}
		return int32(v), nil
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return v, nil
	case int64:
		if v > 2147483647 || v < -2147483648 {
			return 0, errors.New("超出 int32 范围")
		}
		return int32(v), nil
	case uint:
		if v > 2147483647 {
			return 0, errors.New("超出 int32 范围")
		}
		return int32(v), nil
	case uint8:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint32:
		if v > 2147483647 {
			return 0, errors.New("超出 int32 范围")
		}
		return int32(v), nil
	case uint64:
		if v > 2147483647 {
			return 0, errors.New("超出 int32 范围")
		}
		return int32(v), nil
	case float32:
		if v > 2147483647 || v < -2147483648 {
			return 0, errors.New("超出 int32 范围")
		}
		return int32(v), nil
	case float64:
		if v > 2147483647 || v < -2147483648 {
			return 0, errors.New("超出 int32 范围")
		}
		return int32(v), nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		i, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return 0, err
		}
		return int32(i), nil
	default:
		return 0, errors.New("不支持的类型")
	}
}

// ToInt64
//
// @Title ToInt64
// @Description: 将给定的值转换为 int64 类型。
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 15:37:59
// @param value
// @return int64
// @return error
func ToInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		if v > 9223372036854775807 {
			return 0, errors.New("超出 int64 范围")
		}
		return int64(v), nil
	case float32:
		if v > 9223372036854775807 || v < -9223372036854775808 {
			return 0, errors.New("超出 int64 范围")
		}
		return int64(v), nil
	case float64:
		if v > 9223372036854775807 || v < -9223372036854775808 {
			return 0, errors.New("超出 int64 范围")
		}
		return int64(v), nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	default:
		return 0, errors.New("不支持的类型")
	}
}

// ToUint8
//
// @Title ToUint8
// @Description: 将给定的值转换为 uint8 类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 15:49:15
// @param value
// @return uint8
// @return error
func ToUint8(value interface{}) (uint8, error) {
	switch v := value.(type) {
	case int:
		if v > 255 || v < 0 {
			return 0, errors.New("超出uint8范围")
		}
		return uint8(v), nil
	case int8:
		if v < 0 {
			return 0, errors.New("超出uint8范围")
		}
		return uint8(v), nil
	case int16:
		if v > 255 || v < 0 {
			return 0, errors.New("超出uint8范围")
		}
		return uint8(v), nil
	case int32:
		if v > 255 || v < 0 {
			return 0, errors.New("超出uint8范围")
		}
		return uint8(v), nil
	case int64:
		if v > 255 || v < 0 {
			return 0, errors.New("超出uint8范围")
		}
		return uint8(v), nil
	case uint:
		if v > 255 {
			return 0, errors.New("超出uint8范围")
		}
		return uint8(v), nil
	case uint8:
		return v, nil
	case uint16:
		if v > 255 {
			return 0, errors.New("超出uint8范围")
		}
		return uint8(v), nil
	case uint32:
		if v > 255 {
			return 0, errors.New("超出uint8范围")
		}
		return uint8(v), nil
	case uint64:
		if v > 255 {
			return 0, errors.New("超出uint8范围")
		}
		return uint8(v), nil
	case float32:
		if v > 255 || v < 0 {
			return 0, errors.New("超出uint8范围")
		}
		return uint8(v), nil
	case float64:
		if v > 255 || v < 0 {
			return 0, errors.New("超出uint8范围")
		}
		return uint8(v), nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		if v == "" {
			return 0, errors.New("空字符串无法转换为uint8")
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		if i > 255 || i < 0 {
			return 0, errors.New("超出uint8范围")
		}
		return uint8(i), nil
	default:
		return 0, errors.New("不支持的类型")
	}
}

// ToUint16
//
// @Title ToUint16
// @Description: 将给定的值转换为 uint16 类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 15:50:05
// @param value
// @return uint16
// @return error
func ToUint16(value interface{}) (uint16, error) {
	switch v := value.(type) {
	case int:
		if v > 65535 || v < 0 {
			return 0, errors.New("超出uint16范围")
		}
		return uint16(v), nil
	case int8:
		if v < 0 {
			return 0, errors.New("超出uint16范围")
		}
		return uint16(v), nil
	case int16:
		if v < 0 {
			return 0, errors.New("超出uint16范围")
		}
		return uint16(v), nil
	case int32:
		if v > 65535 || v < 0 {
			return 0, errors.New("超出uint16范围")
		}
		return uint16(v), nil
	case int64:
		if v > 65535 || v < 0 {
			return 0, errors.New("超出uint16范围")
		}
		return uint16(v), nil
	case uint:
		if v > 65535 {
			return 0, errors.New("超出uint16范围")
		}
		return uint16(v), nil
	case uint8:
		return uint16(v), nil
	case uint16:
		return v, nil
	case uint32:
		if v > 65535 {
			return 0, errors.New("超出uint16范围")
		}
		return uint16(v), nil
	case uint64:
		if v > 65535 {
			return 0, errors.New("超出uint16范围")
		}
		return uint16(v), nil
	case float32:
		if v > 65535 || v < 0 {
			return 0, errors.New("超出uint16范围")
		}
		return uint16(v), nil
	case float64:
		if v > 65535 || v < 0 {
			return 0, errors.New("超出uint16范围")
		}
		return uint16(v), nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		if v == "" {
			return 0, errors.New("空字符串无法转换为uint16")
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		if i > 65535 || i < 0 {
			return 0, errors.New("超出uint16范围")
		}
		return uint16(i), nil
	default:
		return 0, errors.New("不支持的类型")
	}
}

// ToUint32
//
// @Title ToUint32
// @Description: 将给定的值转换为 uint32 类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 15:50:47
// @param value
// @return uint32
// @return error
func ToUint32(value interface{}) (uint32, error) {
	switch v := value.(type) {
	case int:
		if v > 4294967295 || v < 0 {
			return 0, errors.New("超出uint32范围")
		}
		return uint32(v), nil
	case int8:
		if v < 0 {
			return 0, errors.New("超出uint32范围")
		}
		return uint32(v), nil
	case int16:
		if v < 0 {
			return 0, errors.New("超出uint32范围")
		}
		return uint32(v), nil
	case int32:
		if v < 0 {
			return 0, errors.New("超出uint32范围")
		}
		return uint32(v), nil
	case int64:
		if v > 4294967295 || v < 0 {
			return 0, errors.New("超出uint32范围")
		}
		return uint32(v), nil
	case uint:
		if v > 4294967295 {
			return 0, errors.New("超出uint32范围")
		}
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	case uint16:
		return uint32(v), nil
	case uint32:
		return v, nil
	case uint64:
		if v > 4294967295 {
			return 0, errors.New("超出uint32范围")
		}
		return uint32(v), nil
	case float32:
		if v > 4294967295 || v < 0 {
			return 0, errors.New("超出uint32范围")
		}
		return uint32(v), nil
	case float64:
		if v > 4294967295 || v < 0 {
			return 0, errors.New("超出uint32范围")
		}
		return uint32(v), nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		if v == "" {
			return 0, errors.New("空字符串无法转换为uint32")
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		if i > 4294967295 || i < 0 {
			return 0, errors.New("超出uint32范围")
		}
		return uint32(i), nil
	default:
		return 0, errors.New("不支持的类型")
	}
}

// ToUint64
//
// @Title ToUint64
// @Description: 将给定的值转换为 uint64 类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 16:00:09
// @param value
// @return uint64
// @return error
func ToUint64(value interface{}) (uint64, error) {
	switch v := value.(type) {
	case int:
		if v < 0 {
			return 0, errors.New("超出uint64范围")
		}
		return uint64(v), nil
	case int8:
		if v < 0 {
			return 0, errors.New("超出uint64范围")
		}
		return uint64(v), nil
	case int16:
		if v < 0 {
			return 0, errors.New("超出uint64范围")
		}
		return uint64(v), nil
	case int32:
		if v < 0 {
			return 0, errors.New("超出uint64范围")
		}
		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, errors.New("超出uint64范围")
		}
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	case float32:
		if v < 0 {
			return 0, errors.New("超出uint64范围")
		}
		return uint64(v), nil
	case float64:
		if v < 0 {
			return 0, errors.New("超出uint64范围")
		}
		return uint64(v), nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		if v == "" {
			return 0, errors.New("空字符串无法转换为uint64")
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		if i < 0 {
			return 0, errors.New("超出uint64范围")
		}
		return uint64(i), nil
	default:
		return 0, errors.New("不支持的类型")
	}
}

// ToFloat32
//
// @Title ToFloat32
// @Description: 将给定的值转换为 float32 类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 16:03:21
// @param value
// @return float32
// @return error
func ToFloat32(value interface{}) (float32, error) {
	switch v := value.(type) {
	case int:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case uint:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	case float32:
		return v, nil
	case float64:
		if v > float64(float32(math.MaxFloat32)) || v < float64(float32(-math.MaxFloat32)) {
			return 0, errors.New("超出float32范围")
		}
		return float32(v), nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return 0, err
		}
		return float32(f), nil
	default:
		return 0, errors.New("不支持的类型")
	}
}

// ToFloat64
//
// @Title ToFloat64
// @Description: 将给定的值转换为 float64 类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 16:04:03
// @param value
// @return float64
// @return error
func ToFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
		return f, nil
	default:
		return 0, errors.New("不支持的类型")
	}
}

// ToBool 将字符串或数字转换为 bool
func ToBool(val string) (bool, error) {
	switch val {
	case "1", "true", "True", "TRUE":
		return true, nil
	case "0", "false", "False", "FALSE":
		return false, nil
	default:
		return false, fmt.Errorf("无法将 %s 转换为 bool", val)
	}
}
