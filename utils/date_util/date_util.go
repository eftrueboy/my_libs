package date_util

import "time"

const (
	DATE_SECOND_LAYOUT = "2006-01-02 15:04:05"
)

func Now(layout string) string {
	return time.Now().Format(layout)
}

/**
layout: 2006-01-02 15_04_05
*/
func NowDateSecondFilename() string {
	return Now("2006-01-02 15_04_05")
}

/**
layout: 2006-01-02 15_04_05Z07_00
*/
func NowDateSecondTimezoneFilename() string {
	return Now("2006-01-02 15_04_05Z07_00")
}

/**
layout: 2006-01-02 15_04_05Z07_00
*/
func DateSecondTimezoneFilename(microSeconds int64) string {
	return time.UnixMicro(microSeconds).Format("2006-01-02 15_04_05Z07_00")
}

/**
layout: 2006-01-02 15_04_05
*/
func DateSecondFilename(microSeconds int64) string {
	return time.UnixMicro(microSeconds).Format("2006-01-02 15_04_05")
}

/**
layout: 2006-01-02
*/
func NowDate() string {
	return Now("2006-01-02")
}

/**
layout: 2006-01-02 15
*/
func NowDateHour() string {
	return Now("2006-01-02 15")
}

/**
layout: 2006-01-02 15:04
*/
func NowDateMinute() string {
	return Now("2006-01-02 15:04")
}

/**
layout: 2006-01-02 15:04:05
*/
func NowDateSecond() string {
	return Now("2006-01-02 15:04:05")
}

/**
layout: 2006-01-02 15:04:05.999
*/
func NowDateMillsecond() string {
	return Now("2006-01-02 15:04:05.999")
}

/**
layout: 2006-01-02 15:04:05.999999
*/
func NowDateMicrosecond() string {
	return Now("2006-01-02 15:04:05.999999")
}
