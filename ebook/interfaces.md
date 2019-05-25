## [Go by Example](https://gobyexample.com/): Interfaces

接口是一个方法签名的集合。 所谓方法签名，就是指方法的声明，而不包括实现。

也就是说如果结构体A实现了接口B定义的所有方法，那么A也是B类型。

[interfaces.go](<../src/interfaces.go>)

```go
package main

import "fmt"
import "math"

// 这里定义了一个最基本的表示几何形状的方法的接口
type geometry interface {
    area() float64
    perim() float64
}

// 这里我们要让正方形square和圆形circle实现这个接口
type square struct {
    width, height float64
}
type circle struct {
    radius float64
}

// 在Go中实现一个接口，只要实现该接口定义的所有方法即可
// 下面是正方形实现的接口
func (s square) area() float64 {
    return s.width * s.height
}
func (s square) perim() float64 {
    return 2*s.width + 2*s.height
}

// 圆形实现的接口
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// 如果一个函数的参数是接口类型，那么我们可以使用命名接口
// 来调用这个函数
// 比如这里的正方形square和圆形circle都实现了接口geometry，
// 那么它们都可以作为这个参数为geometry类型的函数的参数。
// 在measure函数内部，Go知道调用哪个结构体实现的接口方法。
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    s := square{width: 3, height: 4}
    c := circle{radius: 5}

    // 这里circle和square都实现了geometry接口，所以
    // circle类型变量和square类型变量都可以作为measure
    // 函数的参数
    measure(s)
    measure(c)
}
```

```bash
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
```

### 空接口

具有零方法的接口称为空接口。它表示为`interface{}`。由于空接口的方法为零，因此所有类型都实现空接口。

如下代码中`describe(i interface{})`函数将空接口作为参数，因此可以传递任何类型。

[interfaces1.go](<../src/interfaces1.go>)

```go
package main

import (  
    "fmt"
)

func describe(i interface{}) {  
    fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {  
    s := "Hello World"
    describe(s)
    i := 55
    describe(i)
    strt := struct {
        name string
    }{
        name: "Naveen R",
    }
    describe(strt)
}
```

```bash
Type = string, value = Hello World
Type = int, value = 55
Type = struct { name string }, value = {Naveen R}
```



类型断言用于提取接口的基础值。

**i.（T）**是用于获取`i`具体类型为的接口的基础值的语法`T`。

```go
v, ok := i.(T)  
```

如果具体类型`i`是`T`，`T`则`v`具有基础值，`i`并且`ok`将为真。

如果具体类型`i`不是`T`那么`ok`将是假的并且`v`将具有类型的零值`T`并且**程序将不会恐慌**。

[interfaces2.go](<../src/interfaces2.go>)

```go
package main

import (  
    "fmt"
)

func assert(i interface{}) {  
    v, ok := i.(int)
    fmt.Println(v, ok)
}
func main() {  
    var s interface{} = 56
    assert(s)
    var i interface{} = "Steven Paul"
    assert(i)
}
```

```bash
56 true
0 false
```

### Type Switch

Type Switch用于将接口的具体类型与各种case语句中指定的多种类型进行比较。它类似于Switch Case。唯一的区别是案例指定类型而不是正常switch中的值。

type switch的语法类似于Type断言。在`i.(T)`Type assertion 的语法中，类型`T`应替换`type`为type switch 的关键字。让我们看看下面的程序如何工作。

[interfaces3.go](<../src/interfaces3.go>)

```go
package main

import (  
    "fmt"
)

func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("I am a string and my value is %s\n", i.(string))
    case int:
        fmt.Printf("I am an int and my value is %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }
}
func main() {  
    findType("Naveen")
    findType(77)
    findType(89.98)
}
```

```bash
I am a string and my value is Naveen
I am an int and my value is 77
Unknown type
```

还可以将类型与接口进行比较。如果我们有一个类型，并且该类型实现了一个接口，则可以将该类型与它实现的接口进行比较。

[interfaces4.go](<../src/interfaces4.go>)

```go
package main

import "fmt"

type Describer interface {  
    Describe()
}
type Person struct {  
    name string
    age  int
}

func (p Person) Describe() {  
    fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {  
    switch v := i.(type) {
    case Describer:
        v.Describe()
    default:
        fmt.Printf("unknown type\n")
    }
}

func main() {  
    findType("Naveen")
    p := Person{
        name: "Naveen R",
        age:  25,
    }
    findType(p)
}
```

```
unknown type
Naveen R is 25 years old
```



### 使用指针接收器和值接收器实现接口

带有值接收器的方法同时接受指针和值接收器。在值为或可以取消引用其值的任何值上调用值方法是合法的。

存储在接口中的具体值是不可寻址的，因此编译器无法自动获取a行号的地址。45因此这段代码失败了。

"d2 = a"如果被取消注释，我们将得到编译错误**.\interfaces5.go:45:8: cannot use a (type Address) as type Describer in assignment:
		Address does not implement Describer (Describe method has pointer receiver)**。这是因为，`Describer`接口是使用第22行中的地址指针接收器实现的，我们正在尝试分配`a`哪个是值类型，并且它没有实现`Describer`接口。

[interfaces5.go](<../src/interfaces5go>)

```go
package main

import "fmt"

type Describer interface {  
    Describe()
}
type Person struct {  
    name string
    age  int
}

func (p Person) Describe() { //implemented using value receiver  
    fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {  
    state   string
    country string
}

func (a *Address) Describe() { //implemented using pointer receiver  
    fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {  
    var d1 Describer
    p1 := Person{"Sam", 25}
    d1 = p1
    d1.Describe()
    p2 := Person{"James", 32}
    d1 = &p2
    d1.Describe()

    var d2 Describer
    a := Address{"Washington", "USA"}

    /* compilation error if the following line is
       uncommented
       cannot use a (type Address) as type Describer
       in assignment: Address does not implement
       Describer (Describe method has pointer
       receiver)
    */
    //d2 = a

    d2 = &a //This works since Describer interface
    //is implemented by Address pointer in line 22
    d2.Describe()

}
```

```
Sam is 25 years old
James is 32 years old
State Washington Country USA
```



### 实现多个接口

一个类型可以实现多个接口。

[interfaces6.go](<../src/interfaces6.go>)

```go
package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    DisplaySalary()
}

type LeaveCalculator interface {  
    CalculateLeavesLeft() int
}

type Employee struct {  
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}

func (e Employee) DisplaySalary() {  
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {  
    return e.totalLeaves - e.leavesTaken
}

func main() {  
    e := Employee {
        firstName: "Naveen",
        lastName: "Ramanathan",
        basicPay: 5000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var s SalaryCalculator = e
    s.DisplaySalary()
    var l LeaveCalculator = e
    fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
}
```

```
Naveen Ramanathan has salary $5200
Leaves left = 25
```

### 嵌入接口

尽管go不提供继承，但可以通过嵌入其他接口来创建新接口。

[interfaces7.go](<../src/interfaces7.go>)

```go
package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    DisplaySalary()
}

type LeaveCalculator interface {  
    CalculateLeavesLeft() int
}

type EmployeeOperations interface {  
    SalaryCalculator
    LeaveCalculator
}

type Employee struct {  
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}

func (e Employee) DisplaySalary() {  
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {  
    return e.totalLeaves - e.leavesTaken
}

func main() {  
    e := Employee {
        firstName: "Naveen",
        lastName: "Ramanathan",
        basicPay: 5000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var empOp EmployeeOperations = e
    empOp.DisplaySalary()
    fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}
```

```
Naveen Ramanathan has salary $5200
Leaves left = 25
```

### 接口零值

接口的零值为零。nil接口既有底层值，也有具体类型为nil。

如果我们尝试在`nil`接口上调用方法，程序将会发生混乱，因为`nil`接口既没有底层值也没有具体类型。

[interfaces8.go](<../src/interfaces8.go>)

```go
package main

import "fmt"

type Describer interface {  
    Describe()
}

func main() {  
    var d1 Describer
    if d1 == nil {
        fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
    }
    d1.Describe()
}
```

```
d1 is nil and has type <nil> value <nil>
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x0 pc=0x492a57]
```



下一篇:[错误 Errors](errors.md)