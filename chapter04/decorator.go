// package strategy
package main

import "fmt"

type Component interface {
	Operation()
}

type ConcreteComponet struct {
}

func (componet *ConcreteComponet) Operation() {
	fmt.Println("this is concrete component")
}

type Decorator struct {
	component Component
}

func NewDecorator(component Component) *Decorator {
	return &Decorator{component: component}
}

func (decorator *Decorator) Operation() {
	decorator.component.Operation()
}

type ConcreteDecorator struct {
	Decorator
}

func NewConcreteDecorator(component Component) *ConcreteDecorator {
	return &ConcreteDecorator{Decorator{component: component}}
}

func (concreteDecorator *ConcreteDecorator) Operation() {
	fmt.Println("this is concrete decorator, ready?")
	concreteDecorator.component.Operation()
}

//func main() {
//	decorator := NewConcreteDecorator(&ConcreteComponet{})
//	decorator.Operation()
//}
