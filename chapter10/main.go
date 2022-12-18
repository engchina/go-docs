package main

import (
	"fmt"
	"time"
)

//example1
//func test() {
//	fmt.Println("I am work in a single goroutine")
//}

//example2
//func printInput(ch chan string) {
//	for val := range ch {
//		if val == "EOF" {
//			break
//		}
//		fmt.Printf("Input is %s\n", val)
//	}
//}

//example3
//func consume(ch chan int) {
//	time.Sleep(time.Second * 100)
//	<-ch
//}

// example4
func send(ch chan int, begin int) {
	for i := begin; i < begin+10; i++ {
		ch <- i
	}
}

func main() {
	//example1
	//go test()
	//time.Sleep(time.Second)

	//example2
	//ch := make(chan string)
	//go printInput(ch)
	//
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	val := scanner.Text()
	//	ch <- val
	//	if val == "EOF" {
	//		fmt.Println("End the game")
	//		break
	//	}
	//}
	//defer close(ch)

	//example3
	//ch := make(chan int, 2)
	//go consume(ch)
	//
	//ch <- 1
	//ch <- 2
	//fmt.Println("I am free")
	//ch <- 2
	//fmt.Println("I can not go there within 100s")
	//time.Sleep(time.Second)
	//
	//defer close(ch)

	//example4
	ch1 := make(chan int)
	ch2 := make(chan int)

	go send(ch1, 0)
	go send(ch2, 10)

	time.Sleep(time.Second)

	for {
		select {
		case val := <-ch1:
			fmt.Printf("get value %d from ch1\n", val)
		case val := <-ch2:
			fmt.Printf("get value %d from ch2\n", val)
		case <-time.After(2 * time.Second):
			fmt.Println("Time out")
			return
		}
	}
}
