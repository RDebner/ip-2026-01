package main

import "fmt"

var base, expo int

func calcularPotencia(a, b int) int {
	var potencia int = a
	for i:= 1; i < b; i++ {
		potencia *= a
	}
	return potencia
}

func main() {
	fmt.Print("Informe o valor da base e do expoente: ")
	fmt.Scan(&base, &expo)
	fmt.Println(calcularPotencia(base, expo))
}