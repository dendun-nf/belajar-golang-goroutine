package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64
	group := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				//x++
				atomic.AddInt64(&x, 1) // no need to worry about race condition
			}
			group.Done()
		}()
	}

	group.Wait()
	//time.Sleep(5 * time.Second)
	fmt.Println("Hasilnya adalah", x)
}
