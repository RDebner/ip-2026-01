package main

import "fmt"

var tempF, tempC, poleg, milim float64

func main() {
	fmt.Print("Escreva a temperatura em Fharenheit: ")
	fmt.Scan(&tempF)
	converterT()
	fmt.Print("Escreva a quantidade de chuva em polegadas: ")
	fmt.Scan(&poleg)
	milim = poleg*25.4
	fmt.Printf("O VALOR EM CELSIUS = %.2f\nA QUANTIDADE DE CHUVA E = %.2f", tempC, milim)
}

func converterT() {
	tempC = 5*(tempF - 32) / 9
}