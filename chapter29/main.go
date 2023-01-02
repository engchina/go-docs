package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}

	wg.Wait()
}
