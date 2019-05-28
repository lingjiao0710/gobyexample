## [Go by Example](https://gobyexample.com/): Channels

Channel是连接并行协程(goroutine)的通道。你可以向一个通道写入数据然后从另外一个通道读取数据。

当我们运行程序的时候，数据ping成功地从一个协程传递到了另外一个协程。 默认情况下，协程之间的通信是同步的，也就是说数据的发送端和接收端必须配对使用。Channel的这种特点使得我们可以不用在程序结尾添加额外的代码也能够获取协程发送端发来的信息。因为程序执行到`msg:=<-messages`的时候被阻塞了，直到获得发送端发来的信息才继续执行。

[channels.go](<../src/channels.go>)

```go
package main

import "fmt"

func main() {

	// 使用`make(chan 数据类型)`来创建一个Channel
	// Channel的类型就是它们所传递的数据的类型
	messages := make(chan string)

	// 使用`channel <-`语法来向一个Channel写入数据
	// 这里我们从一个新的协程向messages通道写入数据ping
	go func() { messages <- "ping" }()

	// 使用`<-channel`语法来从Channel读取数据
	// 这里我们从main函数所在的协程来读取刚刚写入
	// messages通道的数据
	msg := <-messages
	fmt.Println(msg)
}
```

```bash
ping
```

### Deadlock

使用通道时需要考虑的一个重要因素是死锁。如果Goroutine正在通道上发送数据，那么预计其他一些Goroutine应该接收数据。如果没有发生这种情况，程序将在运行时发生恐慌`Deadlock`。

[channels1.go](<../src/channels1.go>)

```go
package main


func main() {  
    ch := make(chan int)
    ch <- 5
}
```

```
fatal error: all goroutines are asleep - deadlock!
```

### 单向通道

可以创建单向通道，即仅发送或接收数据的通道。可以将双向信道转换为仅发送或仅接收信道。

[channels2.go](<../src/channels2.go>)

```go
package main

import (
	"fmt"
)

func sendData(sendch chan<- int){
	sendch <- 10
}

func main() {
	sendch := make(chan int)

	go sendData(sendch)
	fmt.Println(<-sendch)
}
```



### 通道关闭和在通道中使用For range

发送端可以关闭通道以通知接收端不再该通道上发送数据。

接收端可以使用附加变量判断通道是否关闭

```go
v, ok := <- ch  
```

[channels3.go](<../src/channels3.go>)

```go
package main

import (
	"fmt"
)

func producer(chn1 chan int){
	for i := 0; i < 10; i++ {
		chn1 <- i
	}

	close(chn1)
}

func main() {
	ch := make(chan int)
	go producer(ch)

	for{
		v, ok := <-ch
		if ok == false{
			break
		}
		fmt.Println("Received ", v, ok)
	}

	ch1 := make(chan int)
	go producer(ch1)
	for v := range ch1{
		fmt.Println("for range received ", v)
	}

}
```



下一篇:[缓冲通道 Channel Buffering](channelsbuffering.md)