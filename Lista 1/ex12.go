package main

import f "fmt"

var qntHoras, preco int

func main() {
	f.Print("Escreva a quantidade de horas: ")
	f.Scan(&qntHoras)
	preco = ((qntHoras - qntHoras%3)/3)*10 + qntHoras%3*5
	f.Printf("O VALOR A PAGAR E = %.2d", preco)
}