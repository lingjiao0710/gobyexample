package main

import (
	"fmt"
)

func main() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := make([]int, 3, 20)

	var n int
	n = copy(s2, s1)
	fmt.Println(n, s2, len(s2), cap(s2))

	s3 := s1[4:6]
	//s3 cap 从s1[4]开始
	fmt.Println(n, s3, len(s3), cap(s3))

	n = copy(s3, s1[1:5])
	fmt.Println(n, s3, len(s3), cap(s3))
}
