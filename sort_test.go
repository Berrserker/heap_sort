package heap_sor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestRanked struct {
	rank int
}

func (r *TestRanked) Rank() int {
	return r.rank
}

type TestIterable struct {
	ptr  int
	data []*TestRanked
}

func (g *TestIterable) Next() (Ranked, bool) {
	if g.ptr < len(g.data) {
		defer func() { g.ptr++ }()

		return g.data[g.ptr], true
	}

	return nil, false
}

func TestSort(t *testing.T) {
	assert := assert.New(t)

	testArray := make([]*TestRanked, 10)

	for i := range testArray {
		testArray[i] = &TestRanked{rank: i}
	}

	testCase := Sort(&TestIterable{
		ptr:  0,
		data: testArray,
	}, 3)

	assert.Equal(3, len(testCase))
	_, ok := testCase[0].(*TestRanked)
	assert.True(ok)
	assert.ElementsMatch([]int{9, 8, 7}, []int{testCase[0].(*TestRanked).rank, testCase[1].(*TestRanked).rank, testCase[2].(*TestRanked).rank})
}
