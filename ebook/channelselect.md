## [Go by Example](https://gobyexample.com/): Channel Select

Go的select关键字可以让你同时等待多个通道操作，将协程（goroutine），通道（channel）和select结合起来构成了Go的一个强大特性。

[channelselect.go](<../src/channelselect.go>)

```go
package main

import "time"
import "fmt"

func main() {

	// 本例中，我们从两个通道中选择
	c1 := make(chan string)
	c2 := make(chan string)

	// 为了模拟并行协程的阻塞操作，我们让每个通道在一段时间后再写入一个值
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// 我们使用select来等待这两个通道的值，然后输出
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

```

```bash
received one
received two
```

### 默认情况

`select`当其他案例都没有准备就绪时，将执行语句中的默认情况。这通常用于防止select语句阻塞。

[channelselect1.go](<../src/channelselect1.go>)

```go
package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)

	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)

		select {
		case v := <-ch:
			fmt.Println("receive value:", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

}
```

```bash
no value received
no value received
no value received
no value received
no value received
no value received
no value received
no value received
no value received
no value received
receive value: process successful
```

### 死锁

select读取Channel时，当Channel没有Goroutine写入会导致死锁。

```go
package main

func main() {  
    ch := make(chan string)
    select {
    case <-ch:
    }
}
```

```bash
fatal error: all goroutines are asleep - deadlock!
```

如果存在默认情况，则不会发生死锁。

[channelselect2.go](<../src/channelselect2.go>)

```go
package main

import "fmt"

func main() {
	ch := make(chan string)
	select {
	case <-ch:
	default:
		fmt.Println("default case executed")
	}
}
```

```
default case executed
```

### 随机选择

当`select`语句中的多个案例准备就绪时，其中一个案例将随机执行。多次运行程序，结果将不同。

```go
package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	ch <- "from server1"
}

func server2(ch chan string) {
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)
	time.Sleep(time.Second)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}
```

```
from server1
```



下一篇:[超时 timeout](timeout.md)