package carbon

import (
	"time"
)

type Carbon struct {
	baseLoc *time.Location
	abs     bool
	u       time.Time
	t       *Carbon
}

var mill2Seconds = int64(1000)

func NewFromString(layout, t string) Carbon {
	tt, _ := time.Parse(layout, t)

	return Carbon{
		baseLoc: time.UTC,
		abs:     false,
		u:       tt,
	}
}

func NewFromSeconds(sec int64) Carbon {
	tt := time.Unix(sec, 0)

	return Carbon{
		baseLoc: time.UTC,
		abs:     false,
		u:       tt,
	}
}

func NewFromMilliSeconds(msec int64) Carbon {
	tt := time.UnixMilli(msec)

	return Carbon{
		baseLoc: time.UTC,
		abs:     false,
		u:       tt,
	}
}

func NewFromTime(t time.Time) Carbon {
	tt := t

	return Carbon{
		baseLoc: time.UTC,
		abs:     false,
		u:       tt,
	}
}

func (c Carbon) MilliSeconds() int64 {
	return c.u.UnixMilli()
}

func (c Carbon) Seconds() int64 {
	return c.u.Unix()
}

func (c Carbon) SecondNth() int {
	return c.u.Second()
}

func (c Carbon) Minutes() int64 {
	return c.u.Unix() / 60
}

func (c Carbon) MinuteNth() int {
	return c.u.Minute()
}

func (c Carbon) Hours() int64 {
	return c.u.Unix() / 60 * 60
}

func (c Carbon) HourNth() int {
	return c.u.Hour()
}

func (c Carbon) Days() int64 {
	return c.u.Unix() / 60 * 60 * 24
}

func (c Carbon) DayNth() int {
	return c.u.Day()
}

func (c Carbon) Weeks() int64 {
	return c.u.Unix() / 60 * 60 * 24 * 7
}

func (c Carbon) WeekNth() int {
	return int(c.u.Weekday())
}

func (c Carbon) Months() int64 {
	return c.u.Unix() / 60 * 60 * 24 * 30
}

func (c Carbon) MonthNth() int {
	return int(c.u.Month())
}

func (c Carbon) Years() int {
	return c.u.Year()
}

func (c Carbon) Format(layout string) string {
	ll := formatLayout(layout)

	return c.u.Format(ll)
}
