package main

import (
	"fmt"
	"math"
)

var coss, k, diferenca float64

func calFatorial(a float64) float64 {
	for i:=a - 1 ; i > 0; i-- {
		a *= i
	}
	if a == 0 {
		a = 1
	}
	return a
}

func main() {

	fmt.Print("Informe o valor de X: ")
	fmt.Scan(&k)

	coss = 1

	for i:=1; i < 20; i++ {
		if i%2 != 0 {
			coss -= math.Pow(k,float64(i*2))/calFatorial(float64(i*2))
		} else 
		if i%2 == 0 {
			coss += math.Pow(k,float64(i*2))/calFatorial(float64(i*2))
		}
	}

	diferenca = coss - math.Cos(k)

	fmt.Printf("%.2f\n%.2f", coss, diferenca)
}