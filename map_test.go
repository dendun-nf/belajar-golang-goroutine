package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func addToMap(data *sync.Map, group *sync.WaitGroup, num int) {
	defer group.Done()

	data.Store(num, num)
	group.Add(1)
}

func TestMap(t *testing.T) {
	data := new(sync.Map)
	group := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		go addToMap(data, group, i)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
