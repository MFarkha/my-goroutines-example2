package main

import (
	"fmt"
	"sync"
)

func PrintSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}
	var wg sync.WaitGroup
	for i, w := range words {
		wg.Add(1)
		PrintSomething(fmt.Sprintf("%d: %s", i, w), &wg)
	}
	wg.Wait()
}
