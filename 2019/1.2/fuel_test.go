package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuel(t *testing.T) {
	a := assert.New(t)
	a.Equal(fuel(12), 2)
	a.Equal(fuel(14), 2)
	a.Equal(fuel(1969), 654)
	a.Equal(fuel(100756), 33583)
}

func TestStageFuel(t *testing.T) {
	a := assert.New(t)
	a.Equal(stageFuel(14), 2)
	a.Equal(stageFuel(1969), 966)
	a.Equal(stageFuel(100756), 50346)
}
