package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "Hello"
	}()

	data := <-channel
	fmt.Println(data)

	close(channel)
}

func GivemeRespon(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "Hello World 999"
}

func TestGiveMeChannel(t *testing.T) {
	a := make(chan string)

	go GivemeRespon(a)

	data := <-a
	fmt.Println(data)

	close(a)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Hello Dunia"
}

func OnlyOut(channel <-chan string) {
	time.Sleep(1 * time.Second)
	data := <-channel
	fmt.Println(data)
}

func TestInOut(t *testing.T) {
	a := make(chan string)
	go OnlyIn(a)
	go OnlyOut(a)
	time.Sleep(1 * time.Second)
	close(a)
}

func TestBufferedChannel(t *testing.T) {
	a := make(chan string, 3)
	defer close(a)
	go func() {
		a <- "Hello"
		a <- "World"
		a <- "Woi"
	}()

	go func() {
		fmt.Println(<-a)
		fmt.Println(<-a)
		fmt.Println(<-a)
	}()
	fmt.Println("Selesai")
	time.Sleep(5 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		for i := 0; i <= 10; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}
		close(channel)
	}()
	for channel2 := range channel {
		fmt.Println(channel2)
	}
}

func TestSelectChannel(t *testing.T) {
	channel := make(chan string)
	Channel2 := make(chan string)
	defer close(channel)
	defer close(Channel2)
	go GivemeRespon(channel)
	go GivemeRespon(Channel2)
	counter := 0
	for {
		select {
		case data := <-channel:
			fmt.Println("Data dari channel 1:", data)
			counter++
		case data := <-Channel2:
			fmt.Println("Data dari channel 2:", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}
