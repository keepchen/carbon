package carbon

import (
	"fmt"
	"strconv"
	"strings"
)

func floatKeepPrecision(value float64, precision int) float64 {
	builder := strings.Builder{}
	builder.WriteString("%")
	builder.WriteString(".")
	builder.WriteString(fmt.Sprintf("%d", precision))
	builder.WriteString("f")
	layout := builder.String()
	r := fmt.Sprintf(layout, value)
	value, _ = strconv.ParseFloat(r, 64)

	return value
}

func formatLayout(layout string) string {
	ll := strings.ToLower(layout)
	ll = strings.Replace(ll, "y", "2006", 1)
	ll = strings.Replace(ll, "m", "01", 1)
	ll = strings.Replace(ll, "d", "02", 1)
	ll = strings.Replace(ll, "h", "15", 1)
	ll = strings.Replace(ll, "i", "04", 1)
	ll = strings.Replace(ll, "s", "05", 1)
	if strings.Contains(ll, ".sss") {
		ll = strings.Replace(ll, ".sss", ".000", 1)
	}
	if strings.Contains(ll, "sss") {
		ll = strings.Replace(ll, "sss", ".000", 1)
	}

	return ll
}
