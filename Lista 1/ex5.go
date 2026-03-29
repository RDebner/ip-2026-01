package main

import f "fmt"

var numConta int
var tipoConsumidor string
var consumoAgua, preco float64

func main() {
	f.Print("Informe o número da conta, o cunsumo de água e o tipo de consumidor: ")
	f.Scan(&numConta, &consumoAgua, &tipoConsumidor)
	if tipoConsumidor == "R" {
		preco = 5 + consumoAgua*0.05
	} else 
	if tipoConsumidor == "C" {
		preco = 500 + (consumoAgua - 80)*0.25
	} else
	if tipoConsumidor == "I" {
		preco = 800 + (consumoAgua - 100)*0.04
	}

	f.Printf("CONTA = %d\nVALOR DA CONTA = %.2f", numConta, preco)
}