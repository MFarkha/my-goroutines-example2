package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestPrintSomething(t *testing.T) {
	stdOut := os.Stdout // saving stdout to restore it later
	r, w, _ := os.Pipe()
	os.Stdout = w // creating our own os.stdout
	var wg sync.WaitGroup
	wg.Add(1)
	go PrintSomething("sometest", &wg)
	wg.Wait()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "sometest") {
		t.Errorf("Expected to find `sometest`")
	}
}
