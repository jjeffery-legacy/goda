package dt

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// LocalDate represents a date without a time or a timezone.
// Calculations on LocalDate are performed using the standard
// library's time.Time type. For these calculations the time is
// midnight and the timezone is UTC.
type LocalDate struct {
	t time.Time
}

// After reports whether the local date d is after e
func (d LocalDate) After(e LocalDate) bool {
	return d.t.After(e.t)
}

// Before reports whether the local date d is before e
func (d LocalDate) Before(e LocalDate) bool {
	return d.t.Before(e.t)
}

// Equal reports whether d and e represent the same local date.
func (d LocalDate) Equal(e LocalDate) bool {
	return d.t.Equal(e.t)
}

// IsZero reports whether d represents the zero local date,
// January 1, year 1.
func (d LocalDate) IsZero() bool {
	return d.t.IsZero()
}

// Date returns the year, month and day on which d occurs.
func (d LocalDate) Date() (year int, month time.Month, day int) {
	return d.t.Date()
}

// Unix returns d as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC to midnight of the date UTC.
func (d LocalDate) Unix() int64 {
	return d.t.Unix()
}

// Year returns the year in which d occurs.
func (d LocalDate) Year() int {
	return d.t.Year()
}

// Month returns the month of the year specified by d.
func (d LocalDate) Month() time.Month {
	return d.t.Month()
}

// Day returns the day of the month specified by d.
func (d LocalDate) Day() int {
	return d.t.Day()
}

// Weekday returns the day of the week specified by d.
func (d LocalDate) Weekday() time.Weekday {
	return d.t.Weekday()
}

// ISOWeek returns the ISO 8601 year and week number in which d occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to
// week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1
// of year n+1.
func (d LocalDate) ISOWeek() (year, week int) {
	year, week = d.t.ISOWeek()
	return
}

// YearDay returns the day of the year specified by D, in the range [1,365] for non-leap years,
// and [1,366] in leap years.
func (d LocalDate) YearDay() int {
	return d.t.YearDay()
}

const (
	secondsPerDay     = 24 * 60 * 60
	nanosecondsPerDay = secondsPerDay * 1000000000
)

// toDays converts a duration that might contain a fractional number of days
// into an exact number od days. Truncation occurs towards zero. This function
// is used when using durations for date arithmetic.
func toDays(duration time.Duration) time.Duration {
	days := duration.Nanoseconds() / nanosecondsPerDay
	nanoseconds := days * nanosecondsPerDay
	return time.Duration(nanoseconds)

}

// Add returns the local date d + duration.
func (d LocalDate) Add(duration time.Duration) LocalDate {
	t := d.t.Add(toDays(duration))
	return LocalDate{t: t}
}

// Sub returns the duration d-e, which will be an integral number of days.
// If the result exceeds the maximum (or minimum) value that can be stored
// in a Duration, the maximum (or minimum) duration will be returned.
// To compute d-duration, use d.Add(-duration).
func (d LocalDate) Sub(e LocalDate) time.Duration {
	return d.t.Sub(e.t)
}

// AddDate returns the local date corresponding to adding the given number of years,
// months, and days to t. For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
//
// AddDate normalizes its result in the same way that Date does, so, for example,
// adding one month to October 31 yields December 1, the normalized form for November 31.
func (d LocalDate) AddDate(years int, months int, days int) LocalDate {
	t := d.t.AddDate(years, months, days)
	return LocalDate{t: t}
}

// toDate converts the time.Time value into a LocalDate.,
func toLocalDate(t time.Time) LocalDate {
	y, m, d := t.Date()
	return Date(y, m, d)
}

// Today returns the current local date.
func Today() LocalDate {
	return toLocalDate(time.Now())
}

// Date returns the LocalDate corresponding to yyyy-mm-dd.
//
// The month and day values may be outside their usual ranges
// and will be normalized during the conversion.
// For example, October 32 converts to November 1.
func Date(year int, month time.Month, day int) LocalDate {
	return LocalDate{
		t: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}
}

// String returns a string representation of d. The date
// format returned is compatible with ISO 8601: yyyy-mm-dd.
func (d LocalDate) String() string {
	return toString(d)
}

// toString returns the string representation of the date.
func toString(d LocalDate) string {
	year, month, day := d.Date()
	sign := ""
	if year < 0 {
		year = -year
		sign = "-"
	}
	return fmt.Sprintf("%s%04d-%02d-%02d", sign, year, int(month), day)
}

// toQuotedString returns the string representation of the date in quotation marks.
func toQuotedString(d LocalDate) string {
	return fmt.Sprintf(`"%s"`, toString(d))
}

var calendarDateFormats = [...]*regexp.Regexp{
	// ISO 8601 representations
	regexp.MustCompile(`^(-?\d{4})-(\d{1,2})-(\d{1,2})$`),
	regexp.MustCompile(`^(-?\d{4})(\d{2})(\d{2})$`),

	// Not ISO 8601, but still unambiguous
	regexp.MustCompile(`^(-?\d{4})\.(\d{1,2})\.(\d{1,2})$`),
	regexp.MustCompile(`^(-?\d{4})/(\d{1,2})/(\d{1,2})$`),
}

var ordinalDateFormats = [...]*regexp.Regexp{
	// ISO 8601 representations
	regexp.MustCompile(`^(-?\d{4})-(\d{3})$`),
	regexp.MustCompile(`^(-?\d{4})(\d{3})$`),
}

var (
	errInvalidDateFormat = errors.New("invalid date format")
)

// ParseDate attempts to parse a string into a local date. Leading
// and trailing space and quotation marks are ignored. The following
// date formates are recognised: yyyy-mm-dd, yyyymmdd, yyyy.mm.dd,
// yyyy/mm/dd, yyyy-ddd, yyyyddd.
func ParseDate(s string) (LocalDate, error) {
	s = strings.Trim(s, " \t\"'")
	for _, regexp := range calendarDateFormats {
		match := regexp.FindStringSubmatch(s)
		if match != nil {
			// no error checking here because matching the regexp
			// guarantees that parsing the strings will succeed.
			year, _ := strconv.ParseInt(match[1], 10, 0)
			month, _ := strconv.ParseInt(match[2], 10, 0)
			day, _ := strconv.ParseInt(match[3], 10, 0)
			return Date(int(year), time.Month(month), int(day)), nil
		}
	}

	for _, regexp := range ordinalDateFormats {
		match := regexp.FindStringSubmatch(s)
		if match != nil {
			// no error checking here because matching the regexp
			// guarantees that parsing the strings will succeed.
			year, _ := strconv.ParseInt(match[1], 10, 0)
			dayOfYear, _ := strconv.ParseInt(match[2], 10, 0)
			duration := time.Duration((dayOfYear - 1) * nanosecondsPerDay)
			return Date(int(year), 1, 1).Add(duration), nil
		}
	}

	return LocalDate{}, errInvalidDateFormat
}

// MarshalJSON implements the json.Marshaler interface.
// The date is a quoted string in an ISO 8601 format (yyyy-mm-dd).
func (d LocalDate) MarshalJSON() ([]byte, error) {
	return []byte(toQuotedString(d)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The date is expected to be a quoted string in an ISO 8601
// format (calendar or ordinal).
func (d *LocalDate) UnmarshalJSON(data []byte) (err error) {
	s := string(data)
	*d, err = ParseDate(s)
	return
}

// MarshalText implements the encoding.TextMarshaller interface.
// The date format is yyyy-mm-dd.
func (d LocalDate) MarshalText() ([]byte, error) {
	return []byte(toString(d)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaller interface.
// The date is expected to an ISO 8601 format (calendar or ordinal).
func (d *LocalDate) UnmarshalText(data []byte) (err error) {
	s := string(data)
	*d, err = ParseDate(s)
	return
}
