package main

import (
	"os"
	"io"
)

type ioReader interface {
	io.Reader
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
