package carbon

type day struct {
	abs           bool
	opMillSeconds float64
}

type DaySvc interface {
	Abs() DaySvc
	Int64() int64
	Float64(precision ...int) float64
}

func (d differ) Day() DaySvc {
	return &day{
		abs:           false,
		opMillSeconds: float64(d.carbon.u.UnixMilli() - d.carbon.t.u.UnixMilli()),
	}
}

func (v day) Abs() DaySvc {
	v.abs = true

	return v
}

func (v day) Int64() int64 {
	result := int64(v.opMillSeconds / float64(24*60*60*mill2Seconds))
	if v.abs && result < 0 {
		result = 0 - result
	}
	return result
}

func (v day) Float64(precision ...int) float64 {
	result := v.opMillSeconds / float64(24*60*60*mill2Seconds)
	if v.abs && result < 0 {
		result = 0 - result
	}

	if len(precision) > 0 {
		result = floatKeepPrecision(result, precision[0])
	}

	return result
}
