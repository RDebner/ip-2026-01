package main

import "fmt"

var x, s float64

func calcularFatorial(a float64) float64 {
	for i:=a - 1 ; i > 0; i-- {
		a *= i
	}
	return a
}

func main() {
	fmt.Print("Informe o número: ")
	fmt.Scan(&x)

	s = x

	for i:=1 ; i < 20; i++ {
		if i%2 != 0 {
			s -= x/calcularFatorial(float64(i))
		}
		if i%2 == 0 {
			s += x/calcularFatorial(float64(i))
		}
	}

	fmt.Printf("%.2f", s)
}