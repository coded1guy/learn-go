package main

import (
	"fmt"
	fib "learn-go/fibonnaci"
	"learn-go/sqrt"
	wc "learn-go/wordcount"
)

func main() {
	fmt.Println("Select a script to run:")
	fmt.Println("1. square root")
	fmt.Println("2. word counter")
	fmt.Println("3. fibonnaci (1 - 10)")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		sqrt.Run()
	case 2:
		wc.Run()
	case 3:
		fib.Run()
	default:
		fmt.Println("Invalid choice")
	}
}
