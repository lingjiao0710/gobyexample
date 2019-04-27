## [Go by Example](https://gobyexample.com/): Hello World

第一个程序打印经典的"hello world"信息。

运行该程序，将代码保存为hello-world.go并且运行

```bash
go run hello-world.go
```

想要将程序构建为二进制文件，使用go build，然后直接执行构建的二进制文件。

```bash
go build hello-world.go
./hello-world
```

[hello-world.go](<../src/hello-world.go>)

```go
package main
import "fmt"
func main() {
    fmt.Println("hello world")
}
```



下一篇:[值](value.md)