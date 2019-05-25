## [Go by Example](https://gobyexample.com/): Goroutines

goroutine是一个轻量级的线程。Go应用程序通常会同时运行数千个Goroutines。

### Goroutines优于线程的优势

- 与线程相比，Goroutines成本很小。它们的堆栈大小只有几kb，堆栈可以根据应用程序的需要增长和缩小，而在线程的情况下，堆栈大小必须指定并固定。
- Goroutines被多路复用到较少数量的OS线程。程序中可能一个线程有数千个Goroutines。如果该线程中的任何Goroutine阻止说等待用户输入，则创建另一个OS线程并将剩余的Goroutines移动到新的OS线程。所有这些都由运行时处理，我们作为程序员从这些复杂的细节中抽象出来，并给出一个干净的API来处理并发。
- Goroutines使用Channel进行交流。设计通道可防止使用Goroutines访问共享内存时发生竞争条件。通道可以被认为是Goroutines通信的管道。

[goroutines.go](<../src/goroutines.go>)

```go
package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 假设我们有一个函数叫做f(s)
	// 这里我们使用通常的同步调用来调用函数
	f("direct")

	// 为了能够让这个函数以协程(goroutine)方式
	// 运行使用go f(s)
	// 这个协程将和调用它的协程并行执行
	go f("goroutine")

	// 你也可以为匿名函数开启一个协程运行
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 上面的协程在调用之后就异步执行了，所以程序不用等待它们执行完成
	// 就跳到这里来了，下面的Scanln用来从命令行获取一个输入，然后才
	// 让main函数结束
	// 如果没有下面的Scanln语句，程序到这里会直接退出，而上面的协程还
	// 没有来得及执行完，你将无法看到上面两个协程运行的结果
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
```

```bash
direct : 0
direct : 1
direct : 2
goroutine : 0
goroutine : 1
goroutine : 2
going
done
```



### 启动多个Goroutines

[goroutines1.go](<../src/goroutines1.go>)

```go
package main

import (  
    "fmt"
    "time"
)

func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}
func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
func main() {  
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}
```

```
1 a 2 3 b 4	c 5 d e
main terminated
```

![img](https://golangbot.com/content/images/2017/07/Goroutines-explained.png)

下一篇:[频道 Channels](channels.md)