package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 0)
	fmt.Printf("s1 len :%d\n", len(s1))
	test(s1)
	fmt.Printf("after test s1 len :%d\n", len(s1))

	s1 = test1(s1)
	fmt.Println(s1)
}

func test(s []int) {
	s = append(s, 3)
	//因为原来分配的空间不够，所以在另外一个地址又重新分配了空间
	//原始地址的数据没有变
	fmt.Printf("s len :%d\n", len(s))
}

//正确结果
func test1(s []int) []int{
	s = append(s, 3)
	return s
}