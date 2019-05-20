## [Go by Example](https://gobyexample.com/): Functions

函数是Go语言的重要内容。

[functions1.go](<../src/functions1.go>)

```go
package main

import "fmt"

// 这个函数计算两个int型输入数据的和，并返回int型的和
func plus(a int, b int) int {
	// Go需要使用return语句显式地返回值
	return a + b
}

func main() {
	// 函数的调用方式很简单
	// "名称(参数列表)"
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
}

```

```bash
$go run functions1.go
1+2 = 3


```

Go语言内置支持多返回值，这个在Go语言中用的很多，比如一个函数同时返回结果和错误信息。

[functions2.go](<../src/functions2.go>)

```go
package main

import "fmt"

// 这个函数的返回值为两个int
func vals() (int, int) {
	return 3, 7
}

func main() {

	// 获取函数的两个返回值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 如果你只对多个返回值里面的几个感兴趣
	// 可以使用下划线(_)来忽略其他的返回值
	_, c := vals()
	fmt.Println(c)
}
```

```bash
$go run functions2.go
3
7
7
```

支持可变长参数列表的函数可以支持任意个传入参数，比如fmt.Println函数就是一个支持可变长参数列表的函数。

[functions3.go](<../src/functions3.go>)

```go
package main

import "fmt"

// 这个函数可以传入任意数量的整型参数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// 支持可变长参数的函数调用方法和普通函数一样
	// 也支持只有一个参数的情况
	sum(1, 2)
	sum(1, 2, 3)

	// 如果你需要传入的参数在一个切片中，像下面一样
	// "func(slice...)"把切片打散传入
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
```

```bash
$go run functions2.go
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10
```

将切片传递给函数时，需要确认是否会修改切片内容。切片本身会传递给函数，函数不会创建新的切片。

[functions4.go](<../src/functions4.go>)

```go
package main

import (  
    "fmt"
)

func change(s ...string) {  
    s[0] = "Go"
    s = append(s, "playground")
    fmt.Println(s)
}

func main() {  
    welcome := []string{"hello", "world"}
    change(welcome...)
    fmt.Println(welcome)
}
```

```bash
$go run functions4.go
[Go world playground]
[Go world]
```

下一篇:[闭包Closures](Closures.md)