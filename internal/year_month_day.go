package internal

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ymd_DayBits   = 6  // up to 64 days in a month
	ymd_MonthBits = 4  // up to 16 months per year
	ymd_YearBits  = 15 // 32K range; only need -10K to +10K
)

const (
	ymd_DayMask   = (1 << ymd_DayBits) - 1
	ymd_MonthMask = ((1 << ymd_MonthBits) - 1) << ymd_DayBits
)

type YearMonthDay int32

func YearMonthDay_New(year, month, day int) YearMonthDay {
	value := ((year - 1) << (ymd_DayBits + ymd_MonthBits)) |
		((month - 1) << ymd_DayBits) |
		(day - 1)
	return YearMonthDay(value)
}

// YearMonthDay_Parse is used for testing. It parses text in the format YYYY-MM-DD.
// Absolutely no error checking is performed
func YearMonthDay_Parse(text string) YearMonthDay {
	if strings.HasPrefix(text, "-") {
		ymd := YearMonthDay_Parse(text[1:])
		return YearMonthDay_New(-ymd.Year(), ymd.Month(), ymd.Day())
	}
	bits := strings.Split(text, "-")
	year, _ := strconv.Atoi(bits[0])
	month, _ := strconv.Atoi(bits[1])
	day, _ := strconv.Atoi(bits[2])

	return YearMonthDay_New(year, month, day)
}

func (ymd YearMonthDay) Year() int {
	return int((ymd >> (ymd_DayBits + ymd_MonthBits)) + 1)
}

func (ymd YearMonthDay) Month() int {
	return int(((ymd & ymd_MonthMask) >> ymd_DayBits) + 1)
}

func (ymd YearMonthDay) Day() int {
	return int((ymd & ymd_DayMask) + 1)
}

func (ymd YearMonthDay) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", ymd.Year(), ymd.Month(), ymd.Day())
}
