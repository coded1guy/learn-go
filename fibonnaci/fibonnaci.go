package fib

import (
	"fmt"
)

func fibonnaci() func() int {
	var num, fib_1, fib_2 = 0, 0, 1
	return func() int {
		iter := num
		if iter == 0 || iter == 1 {
			num += 1
			return iter
		}
		curr_fib := fib_1 + fib_2
		fib_1 = fib_2
		fib_2 = curr_fib
		num += 1
		return curr_fib
	}
}

func Run() {
	var num int = 0
	fmt.Print("Enter a positive integer: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if num <= 0 {
		fmt.Println("The number must be positive.")
		return
	}

	fib_term := make([]int, num)
	f := fibonnaci()
	for i := 0; i < num; i++ {
		fib_term[i] = f()
	}
	fmt.Printf("The first %d fibonnaci terms are: %v\n", num, fib_term)
}
