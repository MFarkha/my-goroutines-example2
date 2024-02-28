package main

import (
	"fmt"
	"time"
)

func server1(ch chan<- string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from server1"
	}
}

func server2(ch chan<- string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server2"
	}
}

func main() {
	fmt.Println("Select with channels")
	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case s1 := <-channel1:
			fmt.Printf("case 1: %s\n", s1)
		case s2 := <-channel1:
			fmt.Printf("case 2: %s\n", s2)
		case s3 := <-channel2:
			fmt.Printf("case 3: %s\n", s3)
		case s4 := <-channel2:
			fmt.Printf("case 4: %s\n", s4)
		}
	}
}
