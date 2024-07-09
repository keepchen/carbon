package carbon

type hour struct {
	abs           bool
	opMillSeconds float64
}

type HourSvc interface {
	Abs() HourSvc
	Int64() int64
	Float64(precision ...int) float64
}

func (d differ) Hour() HourSvc {
	return &hour{
		abs:           false,
		opMillSeconds: float64(d.carbon.u.UnixMilli() - d.carbon.t.u.UnixMilli()),
	}
}

func (v hour) Abs() HourSvc {
	v.abs = true

	return v
}

func (v hour) Int64() int64 {
	result := int64(v.opMillSeconds / float64(60*60*mill2Seconds))
	if v.abs && result < 0 {
		result = 0 - result
	}
	return result
}

func (v hour) Float64(precision ...int) float64 {
	result := v.opMillSeconds / float64(60*60*mill2Seconds)
	if v.abs && result < 0 {
		result = 0 - result
	}

	if len(precision) > 0 {
		result = floatKeepPrecision(result, precision[0])
	}

	return result
}
