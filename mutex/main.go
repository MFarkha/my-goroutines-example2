package main

import (
	"fmt"
	"sync"
)

// func updateMsg(msg *string, s string, wg *sync.WaitGroup, m *sync.Mutex) {
func updateMsg(msg *string, s string, wg *sync.WaitGroup) {
	defer wg.Done()
	// m.Lock()
	*msg = s
	// m.Unlock()
}

func main() {
	var wg sync.WaitGroup
	// var mutex sync.Mutex

	msg := "Hellow world"
	wg.Add(2)
	// go updateMsg(&msg, "Hello Mars!", &wg, &mutex)
	// go updateMsg(&msg, "Hello universe!", &wg, &mutex)
	go updateMsg(&msg, "Hello Mars!", &wg)
	go updateMsg(&msg, "Hello universe!", &wg)
	wg.Wait()
	fmt.Println(msg)
}
