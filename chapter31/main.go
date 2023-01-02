package main

import "fmt"

func main() {
	var chan1 chan<- int
	chan1 = make(chan int, 3)
	chan1 <- 1

	var chan2 <-chan int
	chan2 = make(chan int, 3)
	m := <-chan2
	fmt.Println(m)
}
