package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiborv/hilbert-proximity/geo"
)

func TestRangeQuery(t *testing.T) {
	assert.Equal(t, len(testTree.RangeQuery(geo.NewPoint("000", "000"), geo.NewPoint("111", "111"))), pointsInserted)
	numOfMatches := 0

	min := geo.NewPoint("100", "000")
	max := geo.NewPoint("111", "011")
	matches := testTree.RangeQuery(min, max)
	for n := testTree.next; n.next != nil; n = n.next {
		if n.isWithinRange(min, max) {
			numOfMatches++
		}
	}
	assert.Equal(t, 84, numOfMatches)

	assert.Equal(t, len(matches)-1, numOfMatches)
}
