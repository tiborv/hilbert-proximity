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
