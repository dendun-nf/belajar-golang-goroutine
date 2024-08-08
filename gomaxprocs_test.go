package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	total_cpu := runtime.NumCPU()
	fmt.Println("total cpu:", total_cpu)

	total_thread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread:", total_thread)

	total_goroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine:", total_goroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	total_cpu := runtime.NumCPU()
	fmt.Println("total cpu:", total_cpu)

	runtime.GOMAXPROCS(20)
	total_thread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread:", total_thread)

	total_goroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine:", total_goroutine)

	group.Wait()
}
