package main

import "fmt"

func main() {
	var a = new([5]int)
	
	test(*a)
	fmt.Println(a, len(a))

	test1(a)
	fmt.Println(a, len(a))	
}

func test(a [5]int) {
	a[1] = 6
}

func test1(a *[5]int) {
	a[1] = 6
}

