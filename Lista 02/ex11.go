package main

import "fmt"

var x int

func main() {
	fmt.Print("Escreva o valor de x: ")
	fmt.Scan(&x)

	if x == 2 {
		fmt.Print("Valor inválido \n")
		return
	}

	var funcao int = 8 / (2 - x)

	fmt.Println(funcao)
}