package main

import (
	"fmt"
	"sync"
)

var (
	buffer [5]int = [5]int{-1, -1, -1, -1, -1}
	index  int
	wg     sync.WaitGroup
	ch     chan int
)

func producer() {
	for n := 0; n < 20; n++ {
		item := n * n
		index = n % 5
		buffer[index] = item
		fmt.Printf("producer %d %d %v\n", index, item, buffer)
		ch <- 1
	}
	wg.Done()
}

func consumer() {
	var item int
	for n := 0; n < 20; n++ {
		<-ch
		index = n % 5
		item = buffer[index]
		buffer[index] = -1
		fmt.Printf("consumer %d %d %v\n", index, item, buffer)
	}
	wg.Done()
}

func main() {
	ch = make(chan int, 1)
	wg.Add(2)
	go consumer()
	go producer()
	wg.Wait()
}
