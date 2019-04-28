## [Go by Example](https://gobyexample.com/): Variables

Go中，变量被显示声明，编译器会检查函数调用的变量类型的正确性。

使用var来声明一个或多个变量。

go会自动推断初始化变量的类型。

没有初始化声明的变量为零值，例如，int的零值是0。

:= 语法 是声明和初始化变量的缩写，例如本例中的 var f string = "short"。

Go的基本类型有：

- bool
- string
- int int8 int16 int32 int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte //unit8别名
- rune //int32的别名 代表一个Unicode码
- float32 float64
- complex64 complex128

[variables.go](<../src/variables.go>)

```go
// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

package main

import "fmt"

func main() {

    // `var` declares 1 or more variables.
    var a = "initial"
    fmt.Println(a)

    // You can declare multiple variables at once.
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go will infer the type of initialized variables.
    var d = true
    fmt.Println(d)

    // Variables declared without a corresponding
    // initialization are _zero-valued_. For example, the
    // zero value for an `int` is `0`.
    var e int
    fmt.Println(e)

    // The `:=` syntax is shorthand for declaring and
    // initializing a variable, e.g. for
    // `var f string = "short"` in this case.
    f := "short"
    fmt.Println(f)
}
```

```bash
$ go run variables.go
initial
1 2
true
0
short


```

下一篇:[常量](constants.md)