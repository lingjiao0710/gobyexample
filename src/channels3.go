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