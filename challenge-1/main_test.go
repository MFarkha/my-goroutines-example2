package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	mockMsg := "Hello someone?"
	var wg sync.WaitGroup
	msg = "Hello first!"
	wg.Add(1)
	updateMessage(mockMsg, &wg)
	wg.Wait()

	if msg != mockMsg {
		t.Errorf("updateMessage: expected '%s', received '%s'", mockMsg, msg)
	}
}

func Test_printMessage(t *testing.T) {
	mockMsg := "Hello everyone!"
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	msg = mockMsg
	printMessage()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	expected := strings.TrimSpace(string(result))
	os.Stdout = stdOut

	if expected != mockMsg {
		t.Errorf("printMessage: expected `%s`, received: `%s`", mockMsg, expected)
	}
}
