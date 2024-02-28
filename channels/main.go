package main

import (
	"fmt"
	"log"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)
	fmt.Println("Type something and press enter, unless you want to quit (type `q`)")
	go shout(ping, pong)

	for {
		var userInput string
		fmt.Print("-> ")
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			log.Fatalf("error scanning stdin: %s", err.Error())
		}
		if strings.ToLower(userInput) == "q" {
			break
		}
		ping <- userInput
		output := <-pong
		fmt.Println(output)
	}
	fmt.Println("All done, closing channels")
	close(ping)
	close(pong)
}
