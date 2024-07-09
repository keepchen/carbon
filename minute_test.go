package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinute_Int64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2024-10-10 12:00:00.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(Minute:Int64)", func(t *testing.T) {
		t.Log(diff.Minute().Int64())
		t.Log(diff.Minute().Abs().Int64())
		assert.GreaterOrEqual(t, diff.Minute().Abs().Int64(), int64(0))
	})
}

func TestMinute_Float64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2024-10-10 12:00:00.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(Minute:Float64)", func(t *testing.T) {
		t.Log(diff.Minute().Float64())
		t.Log(diff.Minute().Abs().Float64())
		assert.GreaterOrEqual(t, diff.Minute().Abs().Float64(), float64(0))
	})
}
