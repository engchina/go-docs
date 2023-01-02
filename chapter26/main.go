package main

import (
	"chapter26/a"
	"chapter26/b"
	"chapter26/c"
	"fmt"
)

var A string = initA()

func initA() string {
	fmt.Println("in main.initA()")
	return "a"
}

func init() {
	fmt.Println("in main.init()")
}

func main() {
	a.InitA()
	c.InitC()
	b.InitB()
	fmt.Println("in main.main()")
}
