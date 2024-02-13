package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// CustomTime 是一个自定义的时间类型，用于对 time.Time 进行扩展
type CustomTime time.Time

// Value
//
// @Title Value
// @Description:  将 CustomTime 转换为数据库驱动程序可以识别的值
//
//	对时间的格式化输出。 年份为4位，月份、日期、小时、分钟、秒钟均为2位的格式
//
// @Author liuxingyu
// @Date 2024-02-13 02:22:47
// @receiver ct
// @return driver.Value
// @return error
func (ct CustomTime) Value() (driver.Value, error) {
	// 将 CustomTime 转换为 time.Time 类型
	t := time.Time(ct)

	// 如果时间为零值，则返回 nil
	if t.IsZero() {
		return nil, nil
	}

	// 将时间格式化为字符串并返回
	return t.Format("2006-01-02 15:04:05"), nil
}

// String
//
// @Title String
// @Description: String 将 CustomTime 格式化为字符串
//
//	对时间的格式化输出。 年份为4位，月份、日期、小时、分钟、秒钟均为2位的格式
//
// @Author liuxingyu
// @Date 2024-02-13 02:24:30
// @receiver ct
// @return string
func (ct CustomTime) String() string {
	// 将 CustomTime 转换为 time.Time 类型
	t := time.Time(ct)

	// 将时间格式化为字符串并返回
	return t.Format("2006-01-02 15:04:05")
}

// Scan
//
// @Title Scan
// @Description: 将数据库中的值扫描到 CustomTime 中
// @Author liuxingyu
// @Date 2024-02-13 02:25:33
// @receiver ct
// @param value
// @return error
func (ct *CustomTime) Scan(value interface{}) error {
	// 如果值为 nil，则将 CustomTime 设置为零值并返回
	if value == nil {
		*ct = CustomTime(time.Time{})
		return nil
	}

	// 将值转换为 time.Time 类型，并将其赋给 CustomTime
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("无法扫描 CustomTime 值")
	}
	*ct = CustomTime(t)
	return nil
}
