## [Go by Example](https://gobyexample.com/): Pointers

Go支持指针，可以用来给函数传递变量的引用。

[pointers.go](<../src/pointers.go>)

```go
package main

import "fmt"

// 我们用两个不同的例子来演示指针的用法
// zeroval函数有一个int类型参数，这个时候传递给函数的是变量的值
func zeroval(ival int) {
	ival = 0
}

// zeroptr函数的参数是int类型指针，这个时候传递给函数的是变量的地址
// 在函数内部对这个地址所指向的变量的任何修改都会反映到原来的变量上。
func zeroptr(iptr *int) {
	*iptr = 0
}

func hello() *int{
	i := 5
	return &i
}

func modify(sls []int){
	sls[0] = 100
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &操作符用来取得i变量的地址
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指针类型也可以输出
	fmt.Println("pointer:", &i)

	//指针的零值是nil
	var b *int
	fmt.Println("b is ", b)

	//可以使用new来创建指针
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
    *size = 85
    fmt.Println("New size value is", *size)

    //可以从函数返回指针,go可以智能地在堆上分配这个函数返回的变量
    d := hello()
    fmt.Println("Value of d ", *d)

    //不要使用指向数组的指针进行函数参数传递，改用切片
    a := [3]int{89, 90, 91}
    modify(a[:])
    fmt.Println(a)

    //指针不支持运算
    //b++ invalid operation: b++


}
```

```bash
$go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0xc000062080
b is  <nil>
Size value is 0, type is *int, address is 0xc0000620c0
New size value is 85
Value of d  5
[100 90 91]
```

下一篇:[结构 Structs](structs.md)