package main

import (
	"os"
)

type File interface {
	Open(name string) (*os.File, error)
	Stat() (os.FileInfo, error)
	Read(b []byte) (n int, err error)
	Close() error
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sumfile(PATH string) string {
	f, err := os.Open(PATH)
	check(err)

	stat, err := f.Stat()
	check(err)

	size := stat.Size()

	bytes := make([]byte, size)
	contents, err := f.Read(bytes)
	check(err)

	return string(bytes[:contents])
}
