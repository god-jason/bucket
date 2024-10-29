package time

import (
	"time"
)

const localDateTimeFormat string = "2006-01-02 15:04:05"

type Time struct {
	time.Time
}

func (t *Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(localDateTimeFormat)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, localDateTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	tm, err := time.ParseInLocation(`"`+localDateTimeFormat+`"`, string(b), time.Local)
	*t = Time{tm}
	return err
}

func (t *Time) String() string {
	return t.Format(localDateTimeFormat)
}

func (t *Time) format() string {
	return t.Format(localDateTimeFormat)
}

func (t *Time) MarshalText() ([]byte, error) {
	return []byte(t.format()), nil
}

func Now() Time {
	return Time{time.Now()}
}

func Parse(str string) (Time, error) {
	tm, err := time.Parse(str, localDateTimeFormat)
	return Time{tm}, err
}

func After(ms int64, fn func()) {
	time.AfterFunc(time.Duration(ms)*time.Millisecond, fn)
}

func Sleep(ms int64, fn func()) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
