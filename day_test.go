package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay_Int64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.123"
		target = "2024-12-21 12:00:00.123"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(Day:Int64)", func(t *testing.T) {
		t.Log(diff.Day().Int64())
		t.Log(diff.Day().Abs().Int64())

		assert.GreaterOrEqual(t, diff.Day().Abs().Int64(), int64(0))
	})
}

func TestDay_Float64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2024-12-21 12:00:00.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(Day:Float64)", func(t *testing.T) {
		t.Log(diff.Day().Float64(3))
		t.Log(diff.Day().Abs().Float64(3))

		assert.GreaterOrEqual(t, diff.Day().Abs().Float64(3), float64(0))
	})
}
