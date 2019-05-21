package main

import "fmt"

func main() {
	add10 := closure(10)//其实是构造了一个加10函数
	fmt.Println(add10(5))
	fmt.Println(add10(6))
	add20 := closure(20)
	fmt.Println(add20(5))
}

func closure(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}

}