package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYear_Int64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2025-11-10 12:00:00.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(Year:Int64)", func(t *testing.T) {
		t.Log(diff.Year().Int64())
		t.Log(diff.Year().Abs().Int64())
		assert.GreaterOrEqual(t, diff.Year().Abs().Int64(), int64(0))
	})
}

func TestYear_Float64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2026-10-10 12:00:00.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(Year:Float64)", func(t *testing.T) {
		t.Log(diff.Year().Float64(3))
		t.Log(diff.Year().Abs().Float64(3))
		assert.GreaterOrEqual(t, diff.Year().Abs().Float64(3), float64(0))
	})
}
