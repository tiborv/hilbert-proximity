package geo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	p := NewPoint("01", "11")
	concat, end := p.GetConcatAt(0)
	assert.Equal(t, concat, byte(1), "Should be 01")
	assert.False(t, end, "Should not be end of coordinates")

	concat, end = p.GetConcatAt(1)
	assert.Equal(t, concat, byte(3), "Should be 11")
	assert.True(t, end, "Should be end of coordinates")

	p2 := NewPoint("01", "11")
	assert.True(t, p.Equals(p2), "Should be equal")
	p3 := NewPoint("10", "11")
	assert.False(t, p.Equals(p3), "Should not be equal")

}
