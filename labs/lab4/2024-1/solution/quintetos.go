package main

import (
	"fmt"
	"sync"
)

// Saymon N. - 20211866

var (
	wg sync.WaitGroup
	c []chan int
	count int 
)

func task(n int) {
	defer wg.Done()
	for x := 0; x < 50; x++ {
		fmt.Printf("%d", n)
		count = count + 1
		if (count%5) == 0 {
			fmt.Println() // print the endl character
			for j:=0; j<5; j++ {
				c[j] <- 1 // continue printing
			}
		}
		<- c[n] // wait to keep printing
	}
}

func main() {
	count = 0
	// slice of channels
	c = make([]chan int, 5)
	for i:= range c {
		c[i] = make(chan int, 1)
	}
	// goroutines each with an id
	for x := 0; x < 5; x++ {
		wg.Add(1)
		go task(x)
	}
	wg.Wait()
	// close channels
	for i:=0; i<5; i++ {
		close(c[i])
	}
}
