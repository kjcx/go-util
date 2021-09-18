package date

import "time"

// StrToDate 字符串日期 转时间类型
func StrToDate(timeStr string)time.Time{
	timeStart,_ := time.ParseInLocation("2006-01-02",timeStr,time.Local)
	return timeStart
}

// StrToTime 字符串转时间类型
func StrToTime(timeStr string)time.Time{
	timeStart,_ := time.ParseInLocation("2006-01-02 15:04:05",timeStr,time.Local)
	return timeStart
}