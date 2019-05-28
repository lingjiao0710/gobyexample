## [Go by Example](https://gobyexample.com/): Channels Buffering

默认情况下，通道是不带缓冲区的。 发送端发送数据，同时必须又接收端相应的接收数据。 
而带缓冲区的通道则允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。

[chbuffering.go](<../src/chbuffering.go>)

```go
package main

import "fmt"

func main() {

	// 这里我们定义了一个可以存储字符串类型的带缓冲通道
	// 缓冲区大小为2
	messages := make(chan string, 2)

	// 因为messages是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	messages <- "buffered"
	messages <- "channel"

	// 然后我们和上面例子一样获取这两个数据
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
```

```bash
buffered
channel
```

[chbuffering1.go](<../src/chbuffering1.go>)

```go
package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Successfully wrote ", i, "to ch")
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go write(ch)

	time.Sleep(2 * time.Second)

	for v := range ch {
		fmt.Println("read value ", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}
```

```
Successfully wrote  0 to ch
Successfully wrote  1 to ch
read value  0 from ch
Successfully wrote  2 to ch
read value  1 from ch
Successfully wrote  3 to ch
read value  2 from ch
Successfully wrote  4 to ch
read value  3 from ch
read value  4 from ch
```

### DeadLock

写入缓冲通道时，写入数据超过容量，并且没有读取时，会出现死锁。

[chbuffering2.go](<../src/chbuffering2.go>)

```go
package main

import (  
    "fmt"
)

func main() {  
    ch := make(chan string, 2)
    ch <- "naveen"
    ch <- "paul"
    ch <- "steve"
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

```bash
fatal error: all goroutines are asleep - deadlock!
```



### 长度和容量

缓冲信道的容量是信道可以容纳的值的数量。这是我们使用`make`函数创建缓冲通道时指定的值。

缓冲通道的长度是当前在其中排队的元素数。

[chbuffering3.go](<../src/chbuffering3.go>)

```go
package main

import (  
    "fmt"
)

func main() {  
    ch := make(chan string, 3)
    ch <- "naveen"
    ch <- "paul"
    fmt.Println("capacity is", cap(ch))
    fmt.Println("length is", len(ch))
    fmt.Println("read value", <-ch)
    fmt.Println("new length is", len(ch))
}
```

```
capacity is 3
length is 2
read value naveen
new length is 1
```

下一篇:[通道方向 Channel Directions](channeldirections.md)