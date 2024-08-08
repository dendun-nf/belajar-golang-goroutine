package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	var x int

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x++
				//there would be a time when each go routine will have
				//defined same value of x = number + 1 several times at the same time
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Hasilnya adalah", x)
}
