package main

import "fmt"

type Observer interface {
	Update()
}

type Subject interface {
	GetState() int
	SetState(int)
	Attach(Observer)
	Notify()
}

type FirstObserver struct {
	name    string
	subject Subject
}

func (observer *FirstObserver) Update() {
	fmt.Printf("this is the first observer, state from subject is %d\n", observer.subject.GetState())
}

func NewFirstObserver(subject Subject) *FirstObserver {
	observer := &FirstObserver{subject: subject}
	subject.Attach(observer)
	return observer
}

type ConcreteSubject struct {
	observers []Observer
	state     int
}

func (subject *ConcreteSubject) GetState() int {
	return subject.state
}

func (subject *ConcreteSubject) SetState(state int) {
	subject.state = state
	subject.Notify()
}

func (subject *ConcreteSubject) Attach(observer Observer) {
	subject.observers = append(subject.observers, observer)
}

func (subject *ConcreteSubject) Notify() {
	for _, observer := range subject.observers {
		observer.Update()
	}
}

func main() {
	subject := &ConcreteSubject{}
	_ = NewFirstObserver(subject)
	_ = NewFirstObserver(subject)

	subject.SetState(0)

	//fmt.Println("after notify")
	//first.Update()
	//second.Update()
}
