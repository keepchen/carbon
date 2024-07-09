package carbon

type year struct {
	abs           bool
	opMillSeconds float64
}

type YearSvc interface {
	Abs() YearSvc
	Int64() int64
	Float64(precision ...int) float64
}

func (d differ) Year() YearSvc {
	return &year{
		abs:           false,
		opMillSeconds: float64(d.carbon.u.UnixMilli() - d.carbon.t.u.UnixMilli()),
	}
}

func (v year) Abs() YearSvc {
	v.abs = true

	return v
}

func (v year) Int64() int64 {
	result := int64(v.opMillSeconds / float64(24*365*60*60*mill2Seconds))
	if v.abs && result < 0 {
		result = 0 - result
	}
	return result
}

func (v year) Float64(precision ...int) float64 {
	result := v.opMillSeconds / float64(24*365*60*60*mill2Seconds)
	if v.abs && result < 0 {
		result = 0 - result
	}

	if len(precision) > 0 {
		result = floatKeepPrecision(result, precision[0])
	}

	return result
}
