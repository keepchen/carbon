package carbon

type differ struct {
	carbon Carbon
}

type Differ interface {
	MilliSecond() MilliSecondSvc
	Second() SecondSvc
	Minute() MinuteSvc
	Hour() HourSvc
	Day() DaySvc
	Week() WeekSvc
	Month() MonthSvc
	Year() YearSvc
}

func (c Carbon) Diff(t Carbon) Differ {
	// make u to UTC
	c.u = c.u.In(c.baseLoc)
	// make t to UTC
	t.u = t.u.In(c.baseLoc)
	c.t = &t

	return differ{carbon: c}
}
