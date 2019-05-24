## [Go by Example](https://gobyexample.com/): Methods

一般的函数定义叫做函数，定义在结构体上面的函数叫做该结构体的方法。

[methods.go](<../src/methods.go>)

```go
package main

import "fmt"

type rect struct {
	width, height int
}

// 这个area方法有一个限定类型*rect，
// 表示这个函数是定义在rect结构体上的方法
func (r *rect) area() int {
	return r.width * r.height
}

// 方法的定义限定类型可以为结构体类型
// 也可以是结构体指针类型
// 区别在于如果限定类型是结构体指针类型
// 那么在该方法内部可以修改结构体成员信息
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// 调用方法
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go语言会自动识别方法调用的参数是结构体变量还是
	// 结构体指针，如果你要修改结构体内部成员值，那么使用
	// 结构体指针作为函数限定类型，也就是说参数若是结构体
	//变量，仅仅会发生值拷贝。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
```

```bash
$go run methods.go
area:  50
perim: 30
area:  50
perim: 30
```



从某种意义上说，方法是函数的“语法糖”。当函数与某个特定的类型绑定，那么它就是一个方法。也证因为如此，我们可以将方法“还原”成函数。

instance.method(args)->(type).func(instance,args)

为了区别这两种方式，官方文档中将左边的称为`Method Value`，右边则是`Method Expression`。Method Value是包装后的状态对象，总是与特定的对象实例关联在一起（类似闭包，拐带私奔），而Method Expression函数将Receiver作为第一个显式参数，调用时需额外传递。

注意：对于Method Expression，T仅拥有T Receiver方法，`*T`拥有（T+`*T`）所有方法。

```go
package main

import (
	"fmt"
)

func main() {
	p := Person{2, "张三"}

	p.test(1)
	var f1 func(int) = p.test
	f1(2)
	Person.test(p, 3)
	var f2 func(Person, int) = Person.test
	f2(p, 4)

}

type Person struct {
	Id   int
	Name string
}

func (this Person) test(x int) {
	fmt.Println("Id:", this.Id, "Name", this.Name)
	fmt.Println("x=", x)
}
```

```bash
$go run methods1.go
Id: 2 Name 张三
x= 1
Id: 2 Name 张三
x= 2
Id: 2 Name 张三
x= 3
Id: 2 Name 张三
x= 4
```

使用匿名字段，实现模拟继承。即可直接访问匿名字段（匿名类型或匿名指针类型）的方法这种行为类似“继承”。访问匿名字段方法时，有隐藏规则，这样我们可以实现override效果。

```go
package main

import (
	"fmt"
)

func main() {
	p := Student{Person{2, "张三"}, 25}
	p.test()
    p.Person.test()

}

type Person struct {
	Id   int
	Name string
}

type Student struct {
	Person
	Score int
}

func (this Person) test() {
	fmt.Println("person test")
}

func (this Student) test() {
	fmt.Println("student test")
}
```

```bash
student test
person test
```

非结构类型的方法：内置类型int创建类型别名，然后创建一个使用此类型别名作为接收器的方法。

```go
package main

import "fmt"

type myInt int

func (a myInt) add(b myInt) myInt {  
    return a + b
}

func main() {  
    num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("Sum is", sum)
}
```

```bash
Sum is 15
```

方法中的值接收器与函数中的值参数

当函数有一个value参数时，它只接受一个value参数。

当方法具有值接收器时，它将接受指针和值接收器。

```go
package main

import (  
    "fmt"
)

type rectangle struct {  
    length int
    width  int
}

func area(r rectangle) {  
    fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {  
    fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func main() {  
    r := rectangle{
        length: 10,
        width:  5,
    }
    area(r)
    r.area()

    p := &r
    /*
       compilation error, cannot use p (type *rectangle) as type rectangle 
       in argument to area  
    */
    //area(p)

    p.area()//calling value receiver with a pointer
}
```

```
Area Function result: 50
Area Method result: 50
Area Method result: 50
```

方法中的指针接收器与函数中的指针参数

具有指针参数的函数将仅接受指针，而具有指针接收器的方法将接受值和指针接收器。

```go
package main

import (  
    "fmt"
)

type rectangle struct {  
    length int
    width  int
}

func perimeter(r *rectangle) {  
    fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {  
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func main() {  
    r := rectangle{
        length: 10,
        width:  5,
    }
    p := &r //pointer to r
    perimeter(p)
    p.perimeter()

    /*
        cannot use r (type rectangle) as type *rectangle in argument to perimeter
    */
    //perimeter(r)

    r.perimeter()//calling pointer receiver with a value

}
```

```bash
perimeter function output: 30
perimeter method output: 30
perimeter method output: 30
```



下一篇:[接口 Interfaces](interfaces.md)