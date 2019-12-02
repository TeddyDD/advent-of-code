package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicPrograms(t *testing.T) {
	assert := assert.New(t)
	examples := map[string][]int{
		"1,0,0,0,99":          []int{2, 0, 0, 0, 99},
		"2,3,0,3,99":          []int{2, 3, 0, 6, 99},
		"2,4,4,5,99,0":        []int{2, 4, 4, 5, 99, 9801},
		"1,1,1,4,99,5,6,0,99": []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
	}
	for prog, resMem := range examples {
		c := &computer{}
		err := c.load(prog)
		assert.Nil(err, "load should not return error")
		assert.Nilf(c.run(), "running '%s' should not crash", prog)
		assert.Equal(resMem, c.mem)
	}
}
