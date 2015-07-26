// The code was adapted from the .NET noda library.
// I don't think it will get used.

package internal

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ymdDayBits   = 6  // up to 64 days in a month
	ymdMonthBits = 4  // up to 16 months per year
	ymdYearBits  = 15 // 32K range; only need -10K to +10K
)

const (
	ymdDayMask   = (1 << ymdDayBits) - 1
	ymdMonthMask = ((1 << ymdMonthBits) - 1) << ymdDayBits
)

// YearMonthDay provides a compact representation of a year, month and day
type YearMonthDay int32

// NewYearMonthDay creates a new YearMonthDay. No error checking is performed.
func NewYearMonthDay(year, month, day int) YearMonthDay {
	value := ((year - 1) << (ymdDayBits + ymdMonthBits)) |
		((month - 1) << ymdDayBits) |
		(day - 1)
	return YearMonthDay(value)
}

// ParseYearMonthDay is used for testing. It parses text in the format YYYY-MM-DD.
// Absolutely no error checking is performed.
func ParseYearMonthDay(text string) YearMonthDay {
	if strings.HasPrefix(text, "-") {
		ymd := ParseYearMonthDay(text[1:])
		return NewYearMonthDay(-ymd.Year(), ymd.Month(), ymd.Day())
	}
	bits := strings.Split(text, "-")
	year, _ := strconv.Atoi(bits[0])
	month, _ := strconv.Atoi(bits[1])
	day, _ := strconv.Atoi(bits[2])

	return NewYearMonthDay(year, month, day)
}

// Year returns the year component of the YearMonthDay in the range [-9999,9999].
func (ymd YearMonthDay) Year() int {
	return int((ymd >> (ymdDayBits + ymdMonthBits)) + 1)
}

// Month returns the month component of the YearMonthDay in the range [1,15]
func (ymd YearMonthDay) Month() int {
	return int(((ymd & ymdMonthMask) >> ymdDayBits) + 1)
}

// Day returns the day component of the YearMonthDay in the range [1,63]
func (ymd YearMonthDay) Day() int {
	return int((ymd & ymdDayMask) + 1)
}

// String returns a text representation of the YearMonthDay in ISO format (YYYY-MM-DD).
func (ymd YearMonthDay) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", ymd.Year(), ymd.Month(), ymd.Day())
}
