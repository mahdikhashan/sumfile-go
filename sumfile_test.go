package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"example.com/sumfile/mocks"
)

// TODO: use mockery and pass mocked io.Read interface
func Test_(t *testing.T) {
	str := "0\n2\n3\n4\n5\n"
	r := new(mocks.IoReader)
	r.
		On("Read", mock.Anything).
		Run(func(args mock.Arguments) {
			bytes := args[0].([]byte)
			copy(bytes[:], str)
		}).
		Return(len(str), nil)

	assert.Equal(t, "0\n1\n3\n4\n5\n", sumfile("tmp.txt"), `should be 1 2 3 4 5`)
}
