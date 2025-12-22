package main

import (
	"fmt"
	"time"
)

type Foo struct {
	wait1 chan struct{}
	wait2 chan struct{}
}

func NewFoo() *Foo {
	wait1 := make(chan struct{})
	wait2 := make(chan struct{})
	return &Foo{wait1: wait1, wait2: wait2}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	f.wait1 <- struct{}{}
}

func (f *Foo) Second(printSecond func()) {
	<-f.wait1
	/// Do not change this line
	printSecond()
	f.wait2 <- struct{}{}
}

func (f *Foo) Third(printThird func()) {
	<-f.wait2
	// Do not change this line
	printThird()
}

func main() {
	foo := NewFoo()

	first := func() {
		fmt.Println("first")
	}
	second := func() {
		fmt.Println("second")
	}
	third := func() {
		fmt.Println("third")
	}

	go foo.First(first)
	go foo.Second(second)
	go foo.Third(third)

	time.Sleep(time.Second * 5)
}
