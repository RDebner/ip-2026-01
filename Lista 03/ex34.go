package main

import "fmt"

var c, d, mmc int

func mdc(a,b int) int {
	for b != 0 {
		a,b = b, a%b
	}
	return a
}

func main() {
	fmt.Print("Informe os números para calcular o MMC: ")
	fmt.Scan(&c, &d)

	mmc = (c*d)/mdc(c,d)

	fmt.Printf("O MMC de %d e %d é: %d\n", c, d, mmc)
}