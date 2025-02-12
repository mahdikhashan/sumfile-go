package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sumfile() {
	data, err := os.ReadFile("numbers.txt")
	check(err)

	fmt.Print(string(data))
}

func main() {
	sumfile()
}
