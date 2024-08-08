package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Timer sekarang", time.Now())
	current := <-timer.C
	fmt.Println("Timer sudah habis", current)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("Timer sekarang", time.Now())
	tick := <-channel
	fmt.Println("Timer sudah habis", tick)
}

func TestAfterFunc(t *testing.T) {
	group := &sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(1*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now())
	group.Wait()

}
