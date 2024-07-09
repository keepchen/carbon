package carbon

type second struct {
	abs           bool
	opMillSeconds float64
}

type SecondSvc interface {
	Abs() SecondSvc
	Int64() int64
	Float64(precision ...int) float64
}

func (d differ) Second() SecondSvc {
	return &second{
		abs:           false,
		opMillSeconds: float64(d.carbon.u.UnixMilli() - d.carbon.t.u.UnixMilli()),
	}
}

func (v second) Abs() SecondSvc {
	v.abs = true

	return v
}

func (v second) Int64() int64 {
	result := int64(v.opMillSeconds / float64(mill2Seconds))
	if v.abs && result < 0 {
		result = 0 - result
	}
	return result
}

func (v second) Float64(precision ...int) float64 {
	result := v.opMillSeconds / float64(mill2Seconds)
	if v.abs && result < 0 {
		result = 0 - result
	}

	if len(precision) > 0 {
		result = floatKeepPrecision(result, precision[0])
	}

	return result
}
