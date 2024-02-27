package main

import (
	"sync"
	"testing"
)

func Test_updateMsg(t *testing.T) {
	msg := "Hello Earth!"
	mockMsg1 := "Hello Jupiter!"
	mockMsg2 := "Hello Saturn!"
	var wg sync.WaitGroup
	wg.Add(2)
	updateMsg(&msg, mockMsg1, &wg)
	updateMsg(&msg, mockMsg2, &wg)
	wg.Wait()
	// if mockMsg1 != msg && mockMsg2 != msg {
	// 	t.Errorf("updateMsg: expected `%s` or `%s`, received `%s`", mockMsg1, mockMsg2, msg)
	// }
	if mockMsg2 != msg {
		t.Errorf("updateMsg: expected `%s`, received `%s`", mockMsg2, msg)
	}
}
