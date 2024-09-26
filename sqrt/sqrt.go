package sqrt

import (
	"fmt"
)

func Sqrt(num float64) float64 {
	const precision float64 = 0.0000000000000001

	var root float64 = num / 2

	for i := 1; i <= 10; i++ {
		//fmt.Println(root)
		nroot := root - ((root*root - num) / (2 * root))

		diff := nroot - root
		if diff < 0 {
			diff = diff * -1
		}
		//fmt.Println(nroot, root, diff)

		if diff <= precision {
			break
		} else {
			root = nroot
		}
	}
	return root
}

func Run() {
	var num float64
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

	fmt.Printf("The square root of %g is: %g\n", num, Sqrt(num))
}
