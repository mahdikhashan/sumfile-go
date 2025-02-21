package main

import (
	"os"
)

type File interface {
	Stat() (os.FileInfo, error)
	Read(b []byte) (n int, err error)
	Close() error
}

type FileOpener interface {
	Open(name string) (File, error)
}

type OSFileOpen struct{}

func (o *OSFileOpen) Open(name string) (File, error) {
	return os.Open(name)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sumfile(openFile FileOpener, PATH string) string {
	f, err := openFile.Open(PATH)
	check(err)

	defer f.Close()

	stat, err := f.Stat()
	check(err)

	size := stat.Size()

	bytes := make([]byte, size)
	contents, err := f.Read(bytes)
	check(err)

	return string(bytes[:contents])
}
