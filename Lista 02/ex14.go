package main

import "fmt"

var valorFabrica int
var arCond, pintMet, vidrEle, dirHid string


func main() {
	fmt.Print("Escreva o valor de fábrica do carro: ")
	fmt.Scan(&valorFabrica)
	fmt.Print("Deseja colocar ar condicionado? (S ou N): ")
	fmt.Scan(&arCond)
	fmt.Print("Deseja fazer a pintura metálica? (S ou N): ")
	fmt.Scan(&pintMet)
	fmt.Print("Deseja colocar vidro elétrico? (S ou N): ")
	fmt.Scan(&vidrEle)
	fmt.Print("Deseja colocar direção hidráulica? (S ou N): ")
	fmt.Scan(&dirHid)

	if arCond == "S" {
		valorFabrica = valorFabrica + 1750
	}
	if pintMet == "S" {
		valorFabrica = valorFabrica + 800
	}
	if vidrEle == "S" {
		valorFabrica = valorFabrica + 1200
	}
	if dirHid == "S" {
		valorFabrica = valorFabrica + 2000
	}

	fmt.Printf("Valor final do carro: %d", +valorFabrica)
}
