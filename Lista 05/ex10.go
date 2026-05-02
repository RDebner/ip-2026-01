package main

import "fmt"

func main() {
	var fibonacci = [50]int{0, 1}

	for i := 2; i < len(fibonacci); i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	fmt.Print(fibonacci)
}