package fib

import "fmt"

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
	f := fibonnaci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
