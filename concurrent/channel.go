package main

import "fmt"

func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, y+x
	}
	close(c)
}
func main() {
	channel := make(chan int, 10)

	go fib(cap(channel), channel)
	for i := range channel {
		fmt.Println(i)
	}
}
