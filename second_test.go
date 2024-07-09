package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecond_Int64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2024-10-10 12:00:00.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(Second:Int64)", func(t *testing.T) {
		t.Log(diff.Second().Int64())
		t.Log(diff.Second().Abs().Int64())
		assert.GreaterOrEqual(t, diff.Second().Abs().Int64(), int64(0))
	})
}

func TestSecond_Float64(t *testing.T) {
	var (
		base   = "2024-10-10 19:00:00.000"
		target = "2024-10-10 12:00:03.000"
	)
	diff := NewFromString(layout, base).Diff(NewFromString(layout, target))

	t.Run("diff(Second:Float64)", func(t *testing.T) {
		t.Log(diff.Second().Float64())
		t.Log(diff.Second().Abs().Float64())
		assert.GreaterOrEqual(t, diff.Second().Abs().Float64(), float64(0))
	})
}
