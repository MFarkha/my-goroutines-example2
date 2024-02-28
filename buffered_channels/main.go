package main

import (
	"fmt"
	"time"
)

func listenToChain(ch <-chan int) {
	for {

		i := <-ch
		fmt.Printf("got %d from a channel\n", i)

		time.Sleep(time.Second * 1)
	}
}

func main() {
	ch := make(chan int, 10)

	go listenToChain(ch)

	for i := 0; i < 100; i++ {
		fmt.Printf("sending %d to channel\n", i)
		ch <- i
		fmt.Printf("sent %d to channel\n", i)
	}

	fmt.Println("Done")
	close(ch)
}
