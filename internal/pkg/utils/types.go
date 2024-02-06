package utils

import "time"

// CustomTime 是一个自定义的时间类型，用于对 time.Time 进行扩展
type CustomTime time.Time

// CustomTimeToString
//
// @Title CustomTimeToString
// @Description: 对时间的格式化输出。 年份为4位，月份、日期、小时、分钟、秒钟均为2位的格式
// @Author liuxingyu
// @Date 2024-02-06 21:14:57
// @receiver ct
// @return string
func (ct CustomTime) CustomTimeToString() string {
	t := time.Time(ct)
	return t.Format("2006-01-02 15:04:05")
}
