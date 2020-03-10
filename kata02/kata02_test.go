package kata02

import "testing"
import "github.com/stretchr/testify/assert"

func RunTestChop(t *testing.T, Chop func(value int, array []int) int) {
	assert := assert.New(t)

	assert.Equal(-1, Chop(3, []int{}))
	assert.Equal(-1, Chop(3, []int{1}))
	assert.Equal(0, Chop(1, []int{1}))

	assert.Equal(0, Chop(1, []int{1, 3, 5}))
	assert.Equal(1, Chop(3, []int{1, 3, 5}))
	assert.Equal(2, Chop(5, []int{1, 3, 5}))
	assert.Equal(-1, Chop(0, []int{1, 3, 5}))
	assert.Equal(-1, Chop(2, []int{1, 3, 5}))
	assert.Equal(-1, Chop(4, []int{1, 3, 5}))
	assert.Equal(-1, Chop(6, []int{1, 3, 5}))

	assert.Equal(0, Chop(1, []int{1, 3, 5, 7}))
	assert.Equal(1, Chop(3, []int{1, 3, 5, 7}))
	assert.Equal(2, Chop(5, []int{1, 3, 5, 7}))
	assert.Equal(3, Chop(7, []int{1, 3, 5, 7}))
	assert.Equal(-1, Chop(0, []int{1, 3, 5, 7}))
	assert.Equal(-1, Chop(2, []int{1, 3, 5, 7}))
	assert.Equal(-1, Chop(4, []int{1, 3, 5, 7}))
	assert.Equal(-1, Chop(6, []int{1, 3, 5, 7}))
	assert.Equal(-1, Chop(8, []int{1, 3, 5, 7}))
}

func TestChopV1(t *testing.T) {
	RunTestChop(t, ChopV1)
}

func TestChopV2(t *testing.T) {
	RunTestChop(t, ChopV2)
}

func TestChopV3(t *testing.T) {
	RunTestChop(t, ChopV3)
}

func TestChopV4(t *testing.T) {
	RunTestChop(t, ChopV4)
}

func TestChopV5(t *testing.T) {
	RunTestChop(t, ChopV5)
}
