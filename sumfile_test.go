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

	tempFile, err := os.CreateTemp("", "mockfile")
	if err != nil {
		t.Fatalf("failed to create temp file %q", err)
	}
	defer os.Remove(tempFile.Name())

	m := NewMockFile(c)

	mOpen := NewMockFileOpener(c)
	mOpen.EXPECT().Open("temp.txt").Return(m, nil).Times(2)

	mInfo := NewMockFileInfo(c)
	mInfo.EXPECT().Size().Return(int64(11)).Times(1)

	m.EXPECT().Stat().Return(mInfo, nil).Times(1)

	m.EXPECT().Read(gomock.Len(11)).DoAndReturn(func(b []byte) (int, error) {
		copy(b, str)
		return 11, nil
	}).Times(1)

	m.EXPECT().Close().Return(nil).Times(1)

	file, err := mOpen.Open("temp.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if file == nil {
		t.Fatalf("Expected a valid file, got nil")
	}

	result := sumfile(mOpen, "temp.txt")

	if result != "0\n2\n3\n4\n5\n\x00" {
		t.Errorf("f %q", result)
	}
}
