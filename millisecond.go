package carbon

type milliSecond struct {
	abs            bool
	opMilliSeconds float64
}

type MilliSecondSvc interface {
	Abs() MilliSecondSvc
	Int64() int64
	Float64(precision ...int) float64
}

func (d differ) MilliSecond() MilliSecondSvc {
	return &milliSecond{
		abs:            false,
		opMilliSeconds: float64(d.carbon.u.UnixMilli() - d.carbon.t.u.UnixMilli()),
	}
}

func (v milliSecond) Abs() MilliSecondSvc {
	v.abs = true

	return v
}

func (v milliSecond) Int64() int64 {
	result := int64(v.opMilliSeconds)
	if v.abs && result < 0 {
		result = 0 - result
	}
	return result
}

func (v milliSecond) Float64(precision ...int) float64 {
	result := v.opMilliSeconds
	if v.abs && result < 0 {
		result = 0 - result
	}

	if len(precision) > 0 {
		result = floatKeepPrecision(result, precision[0])
	}

	return result
}
