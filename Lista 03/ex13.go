package main

import "fmt"

var a, b, h float64

func main() {

	a, b = 1,1

	for {
		if a == 99 {
			break
		}
		h += a/b
		a += 2
		b++
	}

	fmt.Printf("%.2f", h)
}	