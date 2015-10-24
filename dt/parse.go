package dt

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

var parseFormats = struct {
	calendarDates  []string
	ordinalDates   []string
	times          []string
	throwAwayTimes []string
}{
	calendarDates: []string{
		`(-?\d{4})-(\d{1,2})-(\d{1,2})`,
		`^(-?\d{4})(\d{2})(\d{2})`,
		// Not ISO 8601, but still unambiguous
		`(-?\d{4})\.(\d{1,2})\.(\d{1,2})`,
		`(-?\d{4})/(\d{1,2})/(\d{1,2})`,
	},
	ordinalDates: []string{
		`(-?\d{4})-(\d{3})`,
		`(-?\d{4})(\d{3})`,
	},
	times: []string{
		`(\d{1,2}):(\d{1,2}):(\d{1,2})(\.\d*)?`,
		`(\d{1,2}):(\d{1,2})`,
		`(\d{2})(\d{2})(\d{2})(\.\d*)?`,
		`(\d{2})(\d{2})`,
	},
	throwAwayTimes: []string{
		`(T[0-9:.zZ+-]*)?`,
	},
}

var parseRegexp = struct {
	calendarDates     []*regexp.Regexp
	ordinalDates      []*regexp.Regexp
	calendarDateTimes []*regexp.Regexp
	ordinalDateTimes  []*regexp.Regexp
}{}

func init() {
	for _, cd := range parseFormats.calendarDates {
		for _, tat := range parseFormats.throwAwayTimes {
			text := "^" + cd + tat + "$"
			parseRegexp.calendarDates = append(parseRegexp.calendarDates, regexp.MustCompile(text))
		}

		text := "^" + cd + "$"
		parseRegexp.calendarDateTimes = append(parseRegexp.calendarDateTimes, regexp.MustCompile(text))

		for _, tod := range parseFormats.times {
			text = "^" + cd + "T" + tod + "$"
			parseRegexp.calendarDateTimes = append(parseRegexp.calendarDateTimes, regexp.MustCompile(text))
		}
	}

	for _, od := range parseFormats.ordinalDates {
		for _, tat := range parseFormats.throwAwayTimes {
			text := "^" + od + tat + "$"
			parseRegexp.ordinalDates = append(parseRegexp.ordinalDates, regexp.MustCompile(text))
		}

		text := "^" + od + "$"
		parseRegexp.ordinalDateTimes = append(parseRegexp.ordinalDateTimes, regexp.MustCompile(text))

		for _, tod := range parseFormats.times {
			text = "^" + od + "T" + tod + "$"
			parseRegexp.ordinalDateTimes = append(parseRegexp.ordinalDateTimes, regexp.MustCompile(text))
		}
	}
}

// ParseDate attempts to parse a string into a local date. Leading
// and trailing space and quotation marks are ignored. The following
// date formates are recognised: yyyy-mm-dd, yyyymmdd, yyyy.mm.dd,
// yyyy/mm/dd, yyyy-ddd, yyyyddd.
func ParseDate(s string) (LocalDate, error) {
	s = strings.Trim(s, " \t\"'")
	for _, regexp := range parseRegexp.calendarDates {
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

	for _, regexp := range parseRegexp.ordinalDates {
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

// MustParseDate is similar to ParseDate, but instead of returning an error it will
// panic if s is not in one of the expected formats.
func MustParseDate(s string) LocalDate {
	d, err := ParseDate(s)
	if err != nil {
		panic(err.Error())
	}
	return d
}

// ParseDateTime attempts to parse a string into a local date-time. Leading
// and trailing space and quotation marks are ignored. The following
// date formates are recognised: yyyy-mm-dd, yyyymmdd, yyyy.mm.dd,
// yyyy/mm/dd, yyyy-ddd, yyyyddd. The following time formats are recognised:
// HH:MM:SS, HH:MM, HHMMSS, HHMM.
func ParseDateTime(s string) (LocalDateTime, error) {
	s = strings.Trim(s, " \t\"'")
	for _, regexp := range parseRegexp.calendarDateTimes {
		match := regexp.FindStringSubmatch(s)
		if match != nil {
			// no error checking here because matching the regexp
			// guarantees that parsing the strings will succeed.
			year, _ := strconv.ParseInt(match[1], 10, 0)
			month, _ := strconv.ParseInt(match[2], 10, 0)
			day, _ := strconv.ParseInt(match[3], 10, 0)

			var hour, minute, second int64
			if len(match) > 4 {
				hour, _ = strconv.ParseInt(match[4], 10, 0)
			}
			if len(match) > 5 {
				minute, _ = strconv.ParseInt(match[5], 10, 0)
			}
			if len(match) > 6 {
				second, _ = strconv.ParseInt(match[6], 10, 0)
			}

			return DateTime(int(year), time.Month(month), int(day), int(hour), int(minute), int(second)), nil
		}
	}

	for _, regexp := range parseRegexp.ordinalDateTimes {
		match := regexp.FindStringSubmatch(s)
		if match != nil {
			// no error checking here because matching the regexp
			// guarantees that parsing the strings will succeed.
			year, _ := strconv.ParseInt(match[1], 10, 0)
			dayOfYear, _ := strconv.ParseInt(match[2], 10, 0)

			var hour, minute, second int64
			if len(match) > 3 {
				hour, _ = strconv.ParseInt(match[3], 10, 0)
			}
			if len(match) > 4 {
				minute, _ = strconv.ParseInt(match[4], 10, 0)
			}
			if len(match) > 5 {
				second, _ = strconv.ParseInt(match[5], 10, 0)
			}

			duration := time.Duration((dayOfYear - 1) * nanosecondsPerDay)
			return DateTime(int(year), 1, 1, int(hour), int(minute), int(second)).Add(duration), nil
		}
	}

	return LocalDateTime{}, errInvalidDateFormat
}

// MustParseDate is similar to ParseDate, but instead of returning an error it will
// panic if s is not in one of the expected formats.
func MustParseDateTime(s string) LocalDateTime {
	dt, err := ParseDateTime(s)
	if err != nil {
		panic(err.Error())
	}
	return dt
}
