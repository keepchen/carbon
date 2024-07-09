package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHour_Int64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2024-10-14 12:00:00.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(Hour:Int64)", func(t *testing.T) {
		t.Log(diff.Hour().Int64())
		t.Log(diff.Hour().Abs().Int64())
		assert.GreaterOrEqual(t, diff.Hour().Abs().Int64(), int64(0))
	})
}

func TestHour_Float64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2024-10-13 12:00:00.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(Hour:Float64)", func(t *testing.T) {
		t.Log(diff.Hour().Float64())
		t.Log(diff.Hour().Abs().Float64())
		assert.GreaterOrEqual(t, diff.Hour().Abs().Float64(), float64(0))
	})
}
