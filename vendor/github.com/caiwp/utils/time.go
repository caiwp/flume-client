package utils

import (
	"strings"
	"time"
)

// DateTimeFormat : default date time format
const DateTimeFormat = "Y-m-d H:i:s"

// DateFormat : default date format
const DateFormat = "Y-m-d"

// StrToTime : 2017-08-21 Y-m-d return 1503244800, nil
func StrToTime(s, format string) (int64, error) {
	t, err := TimeParse(s, format)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// Date : Y-m-d 1503244800 return 2017-08-21
func Date(format string, i int64) string {
	t := time.Unix(i, 0)
	f := dateFormatReplace(format)
	return t.Format(f)
}

// TimeParse use to date => 2017-08-08 12:12:23 format => Y-m-d H:i:s
func TimeParse(date, format string) (time.Time, error) {
	f := dateFormatReplace(format)
	return time.ParseInLocation(f, date, time.Local)
}

var datePatterns = []string{
	"Y", "2006",
	"m", "01",
	"d", "02",
	"H", "15",
	"i", "04",
	"s", "05",
}

func dateFormatReplace(format string) string {
	r := strings.NewReplacer(datePatterns...)
	return r.Replace(format)
}
