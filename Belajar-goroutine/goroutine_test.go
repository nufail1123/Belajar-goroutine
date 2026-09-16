package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateHelloWorld(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Upsss")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println(number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}
