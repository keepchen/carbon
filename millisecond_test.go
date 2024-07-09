package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMilliSecond_Int64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2024-10-10 12:00:00.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(MilliSecond:Int64)", func(t *testing.T) {
		t.Log(diff.MilliSecond().Int64())
		t.Log(diff.MilliSecond().Abs().Int64())
		assert.GreaterOrEqual(t, diff.MilliSecond().Abs().Int64(), int64(0))
	})
}

func TestMilliSecond_Float64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2024-10-10 12:00:00.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(MilliSecond:Float64)", func(t *testing.T) {
		t.Log(diff.MilliSecond().Float64())
		t.Log(diff.MilliSecond().Abs().Float64())
		assert.GreaterOrEqual(t, diff.MilliSecond().Abs().Float64(), float64(0))
	})
}
