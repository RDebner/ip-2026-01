package main

import f "fmt"


var numeros int
var somaNum float64

func main() {
	f.Print("Escreva um número positivo e maior que 1: ")
	f.Scan(&numeros)
	somaNum = 0
	for i := 1; i <= numeros ; i++ {
		somaNum = somaNum + (1/float64(i))
	}

	f.Printf("%.6f", somaNum)
}