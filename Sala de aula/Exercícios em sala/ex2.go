package main

import "fmt"

var nota1, nota2 float64
var peso1, peso2 float64

func main() {
	fmt.Print("Escreva a nota 1 e seu peso: ")
	fmt.Scan(&nota1, &peso1)
	fmt.Print("Escreva a nota 2 e seu peso: ")
	fmt.Scan(&nota2, &peso2)
	calcularMediaPonderada()
}

func calcularMediaPonderada() {
	var media float64 = ((nota1*peso1) + (nota2*peso2))/(peso1 + peso2)
	fmt.Println(media)
}
