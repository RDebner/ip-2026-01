package main

import "fmt"

var a, b, quociente, resto int

func main() {
	fmt.Print("Informe N1 E N2: ")
	fmt.Scan(&a, &b)

	var c int = 0

	for c < a {
		if (a - c) >= b {
			c += b
			quociente++
		} else {
			break
		}
	}

	resto = a - c

	fmt.Printf("Quociente(%d,%d): %d\nResto: %d\n", a, b, quociente, resto)
}