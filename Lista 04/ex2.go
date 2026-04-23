package main

import "fmt"

var reais = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func soma(r []float64, i int)float64 {
	if len(r) == 1 {
		return r[0]
	}
	return r[i] + soma(r[1:], i)
}

func main() {
	fmt.Print(soma(reais, 0))
}