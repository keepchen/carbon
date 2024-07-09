package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeek_Int64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2025-12-10 12:00:00.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(Week:Int64)", func(t *testing.T) {
		t.Log(diff.Week().Int64())
		t.Log(diff.Week().Abs().Int64())
		assert.GreaterOrEqual(t, diff.Week().Abs().Int64(), int64(0))
	})
}

func TestWeek_Float64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2025-11-10 12:00:00.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(Week:Float64)", func(t *testing.T) {
		t.Log(diff.Week().Float64(3))
		t.Log(diff.Week().Abs().Float64(3))
		assert.GreaterOrEqual(t, diff.Week().Abs().Float64(3), float64(0))
	})
}
