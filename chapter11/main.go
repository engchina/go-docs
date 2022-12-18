package main

import (
	"fmt"
	"strconv"
	"sync"
)

//example1
//func main() {
//	var lock sync.Mutex
//	go func() {
//		lock.Lock()
//		defer lock.Unlock()
//		fmt.Println("func1 get lok at " + time.Now().String())
//		time.Sleep(time.Second)
//		fmt.Println("func1 release lock " + time.Now().String())
//	}()
//
//	time.Sleep(time.Second / 10)
//
//	go func() {
//		lock.Lock()
//		defer lock.Unlock()
//		fmt.Println("func2 get lok at " + time.Now().String())
//		time.Sleep(time.Second)
//		fmt.Println("func2 release lock " + time.Now().String())
//	}()
//
//	time.Sleep(time.Second * 4)
//}

//example2
//var rwLock sync.RWMutex
//
//func main() {
//	for i := 0; i < 5; i++ {
//		go func(i int) {
//			rwLock.RLocker()
//			defer rwLock.RLocker()
//			fmt.Println("read func " + strconv.Itoa(i) + " get rlock at " + time.Now().String())
//		}(i)
//	}
//
//	time.Sleep(time.Second / 10)
//	for i := 0; i < 5; i++ {
//		go func(i int) {
//			rwLock.Lock()
//			defer rwLock.Unlock()
//			fmt.Println("write func " + strconv.Itoa(i) + " get wlock at " + time.Now().String())
//		}(i)
//	}
//
//	time.Sleep(time.Second * 5)
//}

//example3
//func main() {
//	var waitGroup sync.WaitGroup
//	waitGroup.Add(5)
//
//	for i := 0; i < 5; i++ {
//		go func(i int) {
//			fmt.Println("work " + strconv.Itoa(i) + " is done at " + time.Now().String())
//			time.Sleep(time.Second / 10)
//			waitGroup.Done()
//		}(i)
//	}
//	waitGroup.Wait()
//	fmt.Println("all works are done at " + time.Now().String())
//}

// example4
var syncMap sync.Map
var waitGroup sync.WaitGroup

func addNumber(begin int) {
	for i := begin; i < begin+3; i++ {
		syncMap.Store(i, i)
	}
	waitGroup.Done()
}

func main() {
	routineSize := 5
	waitGroup.Add(routineSize)
	for i := 0; i < routineSize; i++ {
		go addNumber(i * 10)
	}
	waitGroup.Wait()

	var size int
	syncMap.Range(func(key, value interface{}) bool {
		size++
		return true
	})

	fmt.Println("syncMap current size is " + strconv.Itoa(size))

	value, ok := syncMap.Load(0)
	if ok {
		fmt.Println("key 0 has value", value, "")
	}

}
