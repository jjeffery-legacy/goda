package dt

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToday(t *testing.T) {
	date := Today()
	now := time.Now()

	d1, m1, y1 := now.Date()
	d2, m2, y2 := date.Date()

	assert.Equal(t, d1, d2)
	assert.Equal(t, m1, m2)
	assert.Equal(t, y1, y2)
}

func TestYears(t *testing.T) {
	for year := -9999; year <= 9999; year++ {
		month := 5
		day := 20

		date := Date(year, time.Month(month), day)
		CheckLocalDate(t, date, year, month, day)
	}
}

func TestMonths(t *testing.T) {
	for month := 1; month <= 12; month++ {
		year := 1969
		day := 12

		date := Date(year, time.Month(month), day)
		CheckLocalDate(t, date, year, month, day)
	}
}

func TestDays(t *testing.T) {
	for day := 1; day <= 31; day++ {
		year := 1969
		month := 1

		date := Date(year, time.Month(month), day)
		CheckLocalDate(t, date, year, month, day)
	}
}

func CheckLocalDate(t *testing.T, date LocalDate, year, month, day int) {

	assert.Equal(t, year, date.Year())
	assert.Equal(t, month, int(date.Month()))
	assert.Equal(t, day, date.Day())

	// Calculate expected text representation
	var text string
	if year < 0 {
		text = fmt.Sprintf("%05d-%02d-%02d", year, month, day)

	} else {
		text = fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	}

	assert.Equal(t, text, date.String())

	if date2, err := ParseDate(text); err != nil || !date.Equal(date2) {
		if err != nil {
			t.Errorf("ParseDate: %s: unexpected error: %v", text, err)
		} else {
			t.Errorf("ParseDate: expected=%v, actual=%v", date, date2)
		}
	}

	// for non-negative years, can check parsing with time package
	if year >= 0 {
		if tm, err := time.Parse("2006-01-02", text); err != nil {
			t.Errorf("time.Parse: unexpected error parsing %s: %v", text, err)
		} else {
			y := tm.Year()
			m := int(tm.Month())
			d := tm.Day()
			if y != year {
				t.Errorf("time.Parse: Year: expected %d, actual %d", year, y)
			}
			if m != month {
				t.Errorf("time.Parse: Month: expected %d, actual %d", month, m)
			}
			if d != day {
				t.Errorf("time.Parse: Day: expected %d, actual %d", day, d)
			}
		}
	}

	// check marshalling and unmarshalling JSON
	data, err := date.MarshalJSON()
	if err != nil {
		t.Errorf("MarshalJSON: %s: unexpected error: %v", text, err)
	} else {
		assert.Equal(t, `"`+text+`"`, string(data))
		var date2 LocalDate
		err = date2.UnmarshalJSON(data)
		if err != nil {
			t.Errorf("UnmarshalJSON: %s: unexpected error: %v", text, err)
		} else {
			if !date.Equal(date2) {
				t.Errorf("UnmarshalJSON: expected %s, actual %s", date.String(), date2.String())
			}
		}
	}

	// check marshalling and unmarshalling text
	data, err = date.MarshalText()
	if err != nil {
		t.Errorf("MarshalText: %s: unexpected error: %v", text, err)
	} else {
		assert.Equal(t, text, string(data))
		var date2 LocalDate
		err = date2.UnmarshalText(data)
		if err != nil {
			t.Errorf("UnmarshalText: %s: unexpected error: %v", text, err)
		} else {
			if !date.Equal(date2) {
				t.Errorf("UnmarshalText: expected %s, actual %s", date.String(), date2.String())
			}
		}
	}
}
