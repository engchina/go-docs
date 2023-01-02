package main

import (
	"fmt"
	"time"
)

func main() {
	var intChan = make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 5)
		intChan <- 15
	}()

	var stringChan = make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		stringChan <- "abc"
	}()

	select {
	case v := <-intChan:
		fmt.Println(v)
	case v := <-stringChan:
		fmt.Println(v)
	default:
		fmt.Println("default")
	}
}
