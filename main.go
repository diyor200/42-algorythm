package main

import (
	"fmt"
	"time"
)

type FooBar struct {
	n          int
	fooChannel chan struct{}
	barChannel chan struct{}
}

func NewFooBar(n int) *FooBar {
	fooChannel := make(chan struct{}, 1)
	barChannel := make(chan struct{}, 1)
	return &FooBar{n: n, fooChannel: fooChannel, barChannel: barChannel}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.fooChannel
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.barChannel <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	fb.fooChannel <- struct{}{}
	for i := 0; i < fb.n; i++ {
		// printBar() outputs "bar". Do not change or remove this line.
		<-fb.barChannel
		printBar()
		fb.fooChannel <- struct{}{}
	}
}

func main() {
	fb := NewFooBar(2)
	go fb.Foo(printfoo)
	go fb.Bar(printbar)

	time.Sleep(time.Second)
}

func printfoo() {
	fmt.Print("foo")
}

func printbar() {
	fmt.Print("bar")
}
