package carbon

type month struct {
	abs           bool
	opMillSeconds float64
}

type MonthSvc interface {
	Abs() MonthSvc
	Int64() int64
	Float64(precision ...int) float64
}

func (d differ) Month() MonthSvc {
	return &month{
		abs:           false,
		opMillSeconds: float64(d.carbon.u.UnixMilli() - d.carbon.t.u.UnixMilli()),
	}
}

func (v month) Abs() MonthSvc {
	v.abs = true

	return v
}

func (v month) Int64() int64 {
	result := int64(v.opMillSeconds / float64(30*24*60*60*mill2Seconds))
	if v.abs && result < 0 {
		result = 0 - result
	}
	return result
}

func (v month) Float64(precision ...int) float64 {
	result := v.opMillSeconds / float64(30*24*60*60*mill2Seconds)
	if v.abs && result < 0 {
		result = 0 - result
	}

	if len(precision) > 0 {
		result = floatKeepPrecision(result, precision[0])
	}

	return result
}
