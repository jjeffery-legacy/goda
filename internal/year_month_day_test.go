package internal_test

import (
	"fmt"
	"testing"

	"github.com/jjeffery/goda/internal"
)

func TestAllDates_YearMonthDay(t *testing.T) {
	for year := -9999; year <= 9999; year++ {
		month := 5
		day := 20

		ymd := internal.NewYearMonthDay(year, month, day)
		CheckYearMonthDay(t, ymd, year, month, day)
	}
}

func TestAllMonths_YearMonthDay(t *testing.T) {
	for month := 1; month < 16; month++ {
		year := 1969
		day := 12

		ymd := internal.NewYearMonthDay(year, month, day)
		CheckYearMonthDay(t, ymd, year, month, day)
	}
}

func TestAllDays_YearMonthDay(t *testing.T) {
	for day := 1; day < 64; day++ {
		year := 2029
		month := 1

		ymd := internal.NewYearMonthDay(year, month, day)
		CheckYearMonthDay(t, ymd, year, month, day)
	}
}

func CheckYearMonthDay(t *testing.T, ymd internal.YearMonthDay, year, month, day int) {
	if ymd.Year() != year {
		t.Errorf("ymd.Year(): expected=%d, actual=%d", year, ymd.Year())
	}
	if ymd.Month() != month {
		t.Errorf("ymd.Month(): expected=%d, actual=%d", month, ymd.Month())
	}
	if ymd.Day() != day {
		t.Errorf("ymd.Day(): expected=%d, actual=%d", day, ymd.Day())
	}

	text := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	if ymd.String() != text {
		t.Errorf("ymd.String(): expected=%s, actual=%s", text, ymd.String())
	}

	ymd2 := internal.ParseYearMonthDay(text)
	if ymd2 != ymd {
		t.Errorf("ymd2: expected=%v, actual=%v", ymd, ymd2)
	}
}
