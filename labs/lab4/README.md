# Concurrency
## Concurrency vs Parallelism
> Concurrency is about **dealing with** lots of things at once.
>
> Parallelism is about **doing** lots of things at once.
> 
> \- Rob Pike
# Go
Also referred to as Golang, it is a expressive, concise, clean, and efficient programming language (at least Google says so).
## Hello, world!
In Go, every program is part of a package, and the `main` package serves as the entry point of a Go application. Here we define a simple program that outputs `Hello, World!`.

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
# Concurrency in Go
We will use Go in this lab due to its **concurrency mechanisms**, which are among its most notable features.
## Goroutines
Goroutines are **lightweight threads** managed by the Go runtime. They are used to **execute functions concurrently**.

```go
go func() {
	fmt.Println("This is running in a goroutine")
}()
```

## Channels
Channels are tools for **communicating between goroutines**. They can be **buffered** or **unbuffered**, which means they can allow the passing of one or more than one values at a time to a channel.

```go
ch := make(chan int)

go func() {
	ch <- 41 // sends the value to the channel
}()
value := <- ch // receives the value from the channel
fmt.Println(value)
```

## Mutexes
Mutexes are **synchronization primitives** used in concurrent programming to ensure that multiple threads or **goroutines do not simultaneously access a shared resource**.

```go
package main
import (
	"fmt"
	"sync"
)

func routine(n int) {
	defer wg.Done()
	fmt.Printf("I am goroutine %d\n", n)
}

var wg sync.WaitGroup

func main() {
	for x := 0; x < 5; x++ {
		wg.Add(1)
		go routine(x)
	}
	wg.Wait()
}
```

---
# References
- [Go Official Documentation](https://go.dev/doc/)

