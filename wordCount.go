package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WordCount(s string) map[string]int {
	slice_s := strings.Fields(s)
	var m = make(map[string]int)

	for _, v := range slice_s {
		if _, exist := m[v]; exist {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	return m
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence and I'll return a map of word count: ")
	input_s, _ := reader.ReadString('\n')
	fmt.Printf("You entered: %s\n", input_s)

	fmt.Printf("Word count of is: %v\n", WordCount(input_s))
}
