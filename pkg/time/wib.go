package time

import (
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

var jkt *time.Location

func init() {
	jkt, _ = time.LoadLocation("Asia/Jakarta")
}

func Now() time.Time {

	return time.Now().In(jkt)
}

func Parse(t time.Time) time.Time {

	return t.In(jkt)
}

func Format(t time.Time) string {
	return t.In(jkt).Format(timeLayout)
}

func IsPast(t time.Time) bool {
	return Parse(t).Before(Now())
}
