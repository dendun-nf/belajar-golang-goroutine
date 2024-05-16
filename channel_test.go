package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Deni Riyanto"
		fmt.Println("Data sent to channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
