package main

import(
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_(t *testing.T) {
	str := "0\n2\n3\n4\n5\n"
	ctrl := gomock.NewController(t)

	m := NewMockioReader(ctrl)

	m.
		EXPECT().
		Read("temp.txt").
		Return(str)
}
