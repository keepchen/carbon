package carbon

import (
	"testing"
	"time"
)

var (
	layout  = "2006-01-02 15:04:05.000"
	timeStr = "2024-10-10 10:10:20.123"
	tim, _  = time.Parse(layout, timeStr)
)

func TestNewFromString(t *testing.T) {
	carb := NewFromString(layout, timeStr)
	t.Run("MillSeconds:", func(t *testing.T) {
		t.Log(carb.MilliSeconds())
	})
	t.Run("Seconds:", func(t *testing.T) {
		t.Log(carb.Seconds())
	})
	t.Run("SecondNth:", func(t *testing.T) {
		t.Log(carb.SecondNth())
	})
	t.Run("Minutes:", func(t *testing.T) {
		t.Log(carb.Minutes())
	})
	t.Run("MinuteNth:", func(t *testing.T) {
		t.Log(carb.MinuteNth())
	})
	t.Run("Days:", func(t *testing.T) {
		t.Log(carb.Days())
	})
	t.Run("DayNth:", func(t *testing.T) {
		t.Log(carb.DayNth())
	})
	t.Run("Weeks:", func(t *testing.T) {
		t.Log(carb.Weeks())
	})
	t.Run("WeekNth:", func(t *testing.T) {
		t.Log(carb.WeekNth())
	})
	t.Run("Months:", func(t *testing.T) {
		t.Log(carb.Months())
	})
	t.Run("MonthNth:", func(t *testing.T) {
		t.Log(carb.MonthNth())
	})
	t.Run("Years:", func(t *testing.T) {
		t.Log(carb.Years())
	})
}

func TestNewFromMillSeconds(t *testing.T) {
	carb := NewFromMilliSeconds(tim.UnixMilli())
	t.Run("MillSeconds:", func(t *testing.T) {
		t.Log(carb.MilliSeconds())
	})
	t.Run("Seconds:", func(t *testing.T) {
		t.Log(carb.Seconds())
	})
	t.Run("SecondNth:", func(t *testing.T) {
		t.Log(carb.SecondNth())
	})
	t.Run("Minutes:", func(t *testing.T) {
		t.Log(carb.Minutes())
	})
	t.Run("MinuteNth:", func(t *testing.T) {
		t.Log(carb.MinuteNth())
	})
	t.Run("Days:", func(t *testing.T) {
		t.Log(carb.Days())
	})
	t.Run("DayNth:", func(t *testing.T) {
		t.Log(carb.DayNth())
	})
	t.Run("Weeks:", func(t *testing.T) {
		t.Log(carb.Weeks())
	})
	t.Run("WeekNth:", func(t *testing.T) {
		t.Log(carb.WeekNth())
	})
	t.Run("Months:", func(t *testing.T) {
		t.Log(carb.Months())
	})
	t.Run("MonthNth:", func(t *testing.T) {
		t.Log(carb.MonthNth())
	})
	t.Run("Years:", func(t *testing.T) {
		t.Log(carb.Years())
	})
}

func TestNewFromSeconds(t *testing.T) {
	carb := NewFromSeconds(tim.Unix())
	t.Run("MillSeconds:", func(t *testing.T) {
		t.Log(carb.MilliSeconds())
	})
	t.Run("Seconds:", func(t *testing.T) {
		t.Log(carb.Seconds())
	})
	t.Run("SecondNth:", func(t *testing.T) {
		t.Log(carb.SecondNth())
	})
	t.Run("Minutes:", func(t *testing.T) {
		t.Log(carb.Minutes())
	})
	t.Run("MinuteNth:", func(t *testing.T) {
		t.Log(carb.MinuteNth())
	})
	t.Run("Days:", func(t *testing.T) {
		t.Log(carb.Days())
	})
	t.Run("DayNth:", func(t *testing.T) {
		t.Log(carb.DayNth())
	})
	t.Run("Weeks:", func(t *testing.T) {
		t.Log(carb.Weeks())
	})
	t.Run("WeekNth:", func(t *testing.T) {
		t.Log(carb.WeekNth())
	})
	t.Run("Months:", func(t *testing.T) {
		t.Log(carb.Months())
	})
	t.Run("MonthNth:", func(t *testing.T) {
		t.Log(carb.MonthNth())
	})
	t.Run("Years:", func(t *testing.T) {
		t.Log(carb.Years())
	})
}

func TestNewFromTime(t *testing.T) {
	carb := NewFromTime(tim)
	t.Run("MilliSeconds:", func(t *testing.T) {
		t.Log(carb.MilliSeconds())
	})
	t.Run("Seconds:", func(t *testing.T) {
		t.Log(carb.Seconds())
	})
	t.Run("SecondNth:", func(t *testing.T) {
		t.Log(carb.SecondNth())
	})
	t.Run("Minutes:", func(t *testing.T) {
		t.Log(carb.Minutes())
	})
	t.Run("MinuteNth:", func(t *testing.T) {
		t.Log(carb.MinuteNth())
	})
	t.Run("Days:", func(t *testing.T) {
		t.Log(carb.Days())
	})
	t.Run("DayNth:", func(t *testing.T) {
		t.Log(carb.DayNth())
	})
	t.Run("Weeks:", func(t *testing.T) {
		t.Log(carb.Weeks())
	})
	t.Run("WeekNth:", func(t *testing.T) {
		t.Log(carb.WeekNth())
	})
	t.Run("Months:", func(t *testing.T) {
		t.Log(carb.Months())
	})
	t.Run("MonthNth:", func(t *testing.T) {
		t.Log(carb.MonthNth())
	})
	t.Run("Years:", func(t *testing.T) {
		t.Log(carb.Years())
	})
}

func TestFormat(t *testing.T) {
	t.Run("Format:", func(t *testing.T) {
		carb := NewFromTime(tim)
		t.Log(carb.Format("Y-m-d H:i:s.sss"))
		t.Log(carb.Format("Y年m月d日 H时i分s秒"))
		t.Log(carb.Format("Y年m月d日 H时i分s秒sss毫秒"))
		t.Log(carb.Format("Y/m/d H/i/s/sss"))
		t.Log(carb.Format("Y_m_d H_i_s_sss"))
		t.Log(carb.Format("YmdHis.sss"))
	})
}
