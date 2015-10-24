package dt

import (
	"encoding/xml"
	"fmt"
	"time"
)

// LocalDateTime represents a date-time without a timezone.
// Calculations on LocalDateTime are performed using the standard
// library's time.Time type. For these calculations the
// timezone is UTC.
//
// LocalDateTime is useful in situations where a date and time
// are specified, without reference to a timezone. Although not
// common, it can be useful. For example, a dose of medication
// may be scheduled for a particular date and time, regardless
// of the timezone that the patient is residing in at the time.
//
// Because LocalDateTime does not specify a unique instant in
// time, it has never been necessary to specify to sub-second
// accuracy. For this reason LocalDateTime only specifies the
// time to second accuracy. In actual fact, LocalDateTime would
// probably be fine if it only specified to minute accuracy.
type LocalDateTime struct {
	t time.Time
}

// After reports whether the local date-time d is after e
func (d LocalDateTime) After(e LocalDateTime) bool {
	return d.t.After(e.t)
}

// Before reports whether the local date-time d is before e
func (d LocalDateTime) Before(e LocalDateTime) bool {
	return d.t.Before(e.t)
}

// Equal reports whether dt and e represent the same local date-time.
func (dt LocalDateTime) Equal(e LocalDateTime) bool {
	return dt.t.Equal(e.t)
}

// IsZero reports whether dt represents the zero local date-time,
// Midnight, January 1, year 1.
func (dt LocalDateTime) IsZero() bool {
	return dt.t.IsZero()
}

// Date returns the year, month and day on which dt occurs.
func (dt LocalDateTime) Date() (year int, month time.Month, day int) {
	return dt.t.Date()
}

// Clock returns the hour, minute and second on which dt occurs.
func (dt LocalDateTime) Clock() (hour int, minute int, second int) {
	hour = dt.Hour()
	minute = dt.Minute()
	second = dt.Second()
	return
}

// DateTime returns the year, month, day, hour minute, second and nanosecond on which dt occurs.
func (dt LocalDateTime) DateTime() (year int, month time.Month, day int, hour int, minute int, second int) {
	year, month, day = dt.t.Date()
	hour, minute, second = dt.Clock()
	return
}

// Unix returns d as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC to midnight of the date-time UTC.
func (dt LocalDateTime) Unix() int64 {
	return dt.t.Unix()
}

// Year returns the year in which dt occurs.
func (dt LocalDateTime) Year() int {
	return dt.t.Year()
}

// Month returns the month of the year specified by dt.
func (dt LocalDateTime) Month() time.Month {
	return dt.t.Month()
}

// Day returns the day of the month specified by dt.
func (dt LocalDateTime) Day() int {
	return dt.t.Day()
}

// Hour returns the hour specified by dt.
func (dt LocalDateTime) Hour() int {
	return dt.t.Hour()
}

// Minute returns the minute specified by dt.
func (dt LocalDateTime) Minute() int {
	return dt.t.Minute()
}

// Second returns the second specified by dt.
func (dt LocalDateTime) Second() int {
	return dt.t.Second()
}

// Weekday returns the day of the week specified by d.
func (d LocalDateTime) Weekday() time.Weekday {
	return d.t.Weekday()
}

// ISOWeek returns the ISO 8601 year and week number in which d occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to
// week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1
// of year n+1.
func (d LocalDateTime) ISOWeek() (year, week int) {
	year, week = d.t.ISOWeek()
	return
}

// YearDay returns the day of the year specified by D, in the range [1,365] for non-leap years,
// and [1,366] in leap years.
func (d LocalDateTime) YearDay() int {
	return d.t.YearDay()
}

// Add returns the local date-time d + duration.
func (dt LocalDateTime) Add(duration time.Duration) LocalDateTime {
	t := dt.t.Add(toSeconds(duration))
	return LocalDateTime{t: t}
}

// Sub returns the duration dt-e, which will be an integral number of seconds.
// If the result exceeds the maximum (or minimum) value that can be stored
// in a Duration, the maximum (or minimum) duration will be returned.
// To compute dt-duration, use dt.Add(-duration).
func (dt LocalDateTime) Sub(e LocalDateTime) time.Duration {
	return dt.t.Sub(e.t)
}

// AddDate returns the local date-time corresponding to adding the given number of years,
// months, and days to t. For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
//
// AddDate normalizes its result in the same way that Date does, so, for example,
// adding one month to October 31 yields December 1, the normalized form for November 31.
func (dt LocalDateTime) AddDate(years int, months int, days int) LocalDateTime {
	t := dt.t.AddDate(years, months, days)
	return LocalDateTime{t: t}
}

// toDate converts the time.Time value into a LocalDateTime.,
func toLocalDateTime(t time.Time) LocalDateTime {
	y, m, d := t.Date()
	hour, minute, second := t.Clock()
	return DateTime(y, m, d, hour, minute, second)
}

// Now returns the current local date-time.
func Now() LocalDateTime {
	return toLocalDateTime(time.Now())
}

// Date returns the LocalDateTime corresponding to yyyy-mm-dd.
//
// The month and day values may be outside their usual ranges
// and will be normalized during the conversion.
// For example, October 32 converts to November 1.
func DateTime(year int, month time.Month, day int, hour int, minute int, second int) LocalDateTime {
	return LocalDateTime{
		t: time.Date(year, month, day, hour, minute, second, 0, time.UTC),
	}
}

// String returns a string representation of d. The date
// format returned is compatible with ISO 8601: yyyy-mm-dd.
func (d LocalDateTime) String() string {
	return localDateTimeString(d)
}

// localDateTimeString returns the string representation of the date.
func localDateTimeString(d LocalDateTime) string {
	year, month, day, hour, minute, second := d.DateTime()
	sign := ""
	if year < 0 {
		year = -year
		sign = "-"
	}
	return fmt.Sprintf("%s%04d-%02d-%02dT%02d:%02d:%02d", sign, year, int(month), day, hour, minute, second)
}

// localDateQuotedString returns the string representation of the date in quotation marks.
func localDateQuotedString(d LocalDateTime) string {
	return fmt.Sprintf(`"%s"`, localDateTimeString(d))
}

// MarshalJSON implements the json.Marshaler interface.
// The date is a quoted string in an ISO 8601 format (yyyy-mm-dd).
func (d LocalDateTime) MarshalJSON() ([]byte, error) {
	return []byte(localDateQuotedString(d)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The date is expected to be a quoted string in an ISO 8601
// format (calendar or ordinal).
func (d *LocalDateTime) UnmarshalJSON(data []byte) (err error) {
	s := string(data)
	*d, err = ParseDateTime(s)
	return
}

// MarshalText implements the encoding.TextMarshaller interface.
// The date format is yyyy-mm-dd.
func (d LocalDateTime) MarshalText() ([]byte, error) {
	return []byte(localDateTimeString(d)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaller interface.
// The date is expected to an ISO 8601 format (calendar or ordinal).
func (d *LocalDateTime) UnmarshalText(data []byte) (err error) {
	s := string(data)
	*d, err = ParseDateTime(s)
	return
}

func (d *LocalDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(localDateTimeString(*d), start)
	return nil
}

func (d *LocalDateTime) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := decoder.DecodeElement(&s, &start); err != nil {
		return err
	}

	if ld, err := ParseDateTime(s); err != nil {
		return err
	} else {
		*d = ld
	}
	return nil
}

func (d *LocalDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: d.String(),
	}, nil
}

func (d *LocalDateTime) UnmarshalXMLAttr(attr xml.Attr) error {
	if ld, err := ParseDateTime(attr.Value); err != nil {
		return err
	} else {
		*d = ld
	}
	return nil
}
