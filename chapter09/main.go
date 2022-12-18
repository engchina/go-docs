package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Person interface {
	SayHello(name string)
	Run() string
}

type Hero struct {
	Name  string
	Age   int
	Speed int
}

func (hero *Hero) SayHello(name string) {
	fmt.Println("Hello " + name + ", I am " + hero.Name)
}

func (hero *Hero) Run() string {
	fmt.Println("I am running at speed " + strconv.Itoa(hero.Speed))
	return "Running"
}

func main() {
	var person Person = &Hero{
		Name:  "John",
		Speed: 100,
	}
	valueOfPerson := reflect.ValueOf(person)
	sayHelloMethod := valueOfPerson.MethodByName("SayHello")
	sayHelloMethod.Call([]reflect.Value{reflect.ValueOf("SpiderMan")})
	runMethod := valueOfPerson.MethodByName("Run")
	result := runMethod.Call([]reflect.Value{})
	fmt.Printf("result of run method is %s.", result[0])
}
