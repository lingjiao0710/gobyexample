## [Go by Example](https://gobyexample.com/): Value

Go有各种值类型，包括strings,integers,floats,booleans,等等.

strings可以使用+进行连接。

[value.go](<../src/value.go>)

```go
// Go has various value types including strings,
// integers, floats, booleans, etc. Here are a few
// basic examples.

package main

import "fmt"

func main() {

    // Strings, which can be added together with `+`.
    fmt.Println("go" + "lang")

    // Integers and floats.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Booleans, with boolean operators as you'd expect.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}

```

```bash
$ go run value.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false

```

下一篇:[变量](variables.md)