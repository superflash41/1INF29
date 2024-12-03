// superflash41
package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	c []chan int
	aorb int
)

func worker1() { // c[0]
	for {
		<- c[0]
		fmt.Printf("A")
		if aorb == 1 {
			c[2] <- 1
		} else {
			c[1] <- 1
		}
	}
	wg.Done()
}

func worker2() { // c[1]
	for {
		<- c[1]
		fmt.Printf("B")
		if aorb == 1 {
			c[0] <- 1
		} else {
			c[2] <- 1
		}
	}
	wg.Done()
}

func worker3() { // c[2]
	for {
		<- c[2]
		fmt.Printf("C")
		if aorb == 0 {
			aorb = 1
		} else {
			aorb = 0
		}
		c[3] <- 1
	}
	wg.Done()
}

func worker4() { // c[3]
	for {
		<- c[3]
		fmt.Printf("D")
		c[4] <- 1
	}
	wg.Done()
}

func worker5() { // c[4]
	for {
		<- c[4]
		fmt.Printf("E")
		if aorb == 1 {
			c[1] <- 1
		} else {
			c[0] <- 1
		}
	}
	wg.Done()
}

func main() {
	aorb = 0 // AB:0   BA:1
	// create the slice of channels
	c = make([]chan int, 5)
	for i:= range c {
		c[i] = make(chan int, 1)
	}
	wg.Add(5)
	go worker1()
	go worker2()
	go worker3()
	go worker4()
	go worker5()
	c[0] <- 1 // start with A
	wg.Wait()
	fmt.Printf("\n")
}
