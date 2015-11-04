package geo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Point(t *testing.T) {
	p := NewPoint("01", "11")
	concat, end := p.GetConcatAt(0)
	assert.Equal(t, concat, "01", "Should 01")
	assert.False(t, end, "Should not be end of coordinates")
	assert.Equal(t, p.GetX(), "01")
	assert.Equal(t, p.GetY(), "11")

	concat, end = p.GetConcatAt(1)
	assert.Equal(t, concat, "11", "Should 11")
	assert.True(t, end, "Should be end of coordinates")

	p2 := NewPoint("01", "11")
	assert.True(t, p.Equals(p2), "Should be equal")
	p3 := NewPoint("10", "11")
	assert.False(t, p.Equals(p3), "Should not be equal")

}

func Test_Within(t *testing.T) {
	p1 := NewPoint("1111", "1111")
	assert.Equal(t, int64(15), p1.xInt)
	assert.Equal(t, int64(15), p1.yInt)
	p2 := NewPoint("0011", "0100")
	p3 := NewPoint("0110", "0110")
	p4 := NewPoint("1001", "0110")
	min := NewPoint("0011", "0010")
	max := NewPoint("0111", "0111")

	assert.True(t, p2.WithinArea(min, max))
	assert.True(t, p3.WithinArea(min, max))
	assert.False(t, p4.WithinArea(min, max))
	assert.False(t, p1.WithinArea(min, max))

}
