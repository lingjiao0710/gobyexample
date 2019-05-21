package main

import "fmt"

func main() {

	var fs []func() int

	for i := 0; i < 3; i++ {
		fs = append(fs, func() int {
			return i
		})
	}

	for _, f := range fs {
		fmt.Printf("%p = %v\n", f, f())
	}
}