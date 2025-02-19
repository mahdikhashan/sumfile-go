package main

import(
	"os"
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_(t *testing.T) {
	str := "0\n2\n3\n4\n5\n"

	c := gomock.NewController(t)
	defer c.Finish()

	openFile = func(name string) (*os.File, error) {
		return nil, nil
	}

	m := NewMockFile(c)
	m.EXPECT().Open("temp.txt").Return(m, nil)

	mInfo := NewMockFileInfo(c)
	mInfo.EXPECT().Size().Return(int64(11))

	m.EXPECT().Stat().Return(mInfo, nil)

	m.EXPECT().Read(gomock.Any()).DoAndReturn(func(b []byte) (int, error) {
		copy(b, str)
		return 11, nil
	})
	m.EXPECT().Close().Return(nil)

	result := sumfile("temp.txt")

	expected := "0\n1\n2\n3\n4\n5"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
