package main

import (
	"fmt"
	"time"
)

type ZeroEvenOdd struct {
	n          int
	zeroSignal chan int8
	evenSignal chan int
	oddSignal  chan int
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeroSignal := make(chan int8, 1)
	evenSignal := make(chan int)
	oddSignal := make(chan int)

	zeo := &ZeroEvenOdd{
		n:          n,
		zeroSignal: zeroSignal,
		evenSignal: evenSignal,
		oddSignal:  oddSignal,
	}

	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		<-z.zeroSignal
		printNumber(0)

		if i%2 == 1 {
			z.oddSignal <- i
		} else {
			z.evenSignal <- i
		}

	}
	close(z.evenSignal)
	close(z.oddSignal)
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	z.zeroSignal <- 0
	for v := range z.evenSignal {
		printNumber(v)
		z.zeroSignal <- 0
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for v := range z.oddSignal {
		printNumber(v)
		z.zeroSignal <- 0
	}
}

func main() {
	zeo := NewZeroEvenOdd(5)
	go zeo.Zero(printNumber)
	go zeo.Even(printNumber)
	go zeo.Odd(printNumber)

	time.Sleep(time.Second)
}

func printNumber(n int) {
	fmt.Print(n)
}
