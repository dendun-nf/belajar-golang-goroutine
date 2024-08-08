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

func giveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Deni Riyanto"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go giveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Deni Riyanto"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Deni"
		channel <- "Riyanto"
	}()

	go func() {
		fmt.Println(len(channel))
		fmt.Println(cap(channel))

		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(5 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		channel <- "Deni"
		channel <- "Riyanto"
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Finish")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go giveMeResponse(channel1)
	go giveMeResponse(channel2)

	var counter int
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

	fmt.Println("Finish")
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go giveMeResponse(channel1)
	go giveMeResponse(channel2)

	var counter int
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2", data)
			counter++
		default:
			fmt.Println("Waiting...")
		}

		if counter == 2 {
			break
		}

	}

	//fmt.Println("Finish")
}
