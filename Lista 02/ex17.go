package main

import f "fmt"

var tipoConsumidor string
var consumoAgua, preco float64

func main() {
	f.Print("Informe o tipo de consumidor e o cunsumo de água: ")
	f.Scan(&tipoConsumidor, &consumoAgua)
	if tipoConsumidor == "Residencial" {
		preco = 5 + consumoAgua*0.05
	} else 
	if tipoConsumidor == "Comercial" {
		preco = 500 + (consumoAgua - 80)*0.25
	} else
	if tipoConsumidor == "Industrial" {
		preco = 800 + (consumoAgua - 100)*0.04
	}

	f.Printf("Conta %s. Valor: %.2f", tipoConsumidor, preco)
}