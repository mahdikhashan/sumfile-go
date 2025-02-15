package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: use a mockery and pass mocked io.Read interface
func Test_(t *testing.T) {
	assert.Equal(t, "0\n2\n3\n4\n5\n", sumfile("numbers.txt"), `should be 1 2 3 4 5`)
}
