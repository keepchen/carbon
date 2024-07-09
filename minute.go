package carbon

type minute struct {
	abs           bool
	opMillSeconds float64
}

type MinuteSvc interface {
	Abs() MinuteSvc
	Int64() int64
	Float64(precision ...int) float64
}

func (d differ) Minute() MinuteSvc {
	return &minute{
		abs:           false,
		opMillSeconds: float64(d.carbon.u.UnixMilli() - d.carbon.t.u.UnixMilli()),
	}
}

func (v minute) Abs() MinuteSvc {
	v.abs = true

	return v
}

func (v minute) Int64() int64 {
	result := int64(v.opMillSeconds / float64(60*mill2Seconds))
	if v.abs && result < 0 {
		result = 0 - result
	}
	return result
}

func (v minute) Float64(precision ...int) float64 {
	result := v.opMillSeconds / float64(60*mill2Seconds)
	if v.abs && result < 0 {
		result = 0 - result
	}

	if len(precision) > 0 {
		result = floatKeepPrecision(result, precision[0])
	}

	return result
}
