package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)

	fmt.Println("Hello, Im Running")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Done")
}
