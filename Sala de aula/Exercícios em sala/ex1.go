package main

import "fmt"

func main() {

	var lado1, lado2, lado3 int

	fmt.Print("Informe os lados do triângulo: ")
	fmt.Scan(&lado1, &lado2, &lado3)
	fmt.Println(lado1, lado2, lado3)

	if lado1 < lado2+lado3 || lado2 < lado1+lado3 || lado3 < lado1+lado2 {
		if lado1 == lado2 && lado1 == lado3 && lado2 == lado3 {
			fmt.Print("O triângulo é equilátero.")
		} else if lado1 != lado2 && lado1 != lado3 && lado2 != lado3 {
			fmt.Print("O triângulo é escaleno")
		} else if lado1 == lado2 && lado3 != lado1 || lado1 == lado3 && lado2 != lado1 || lado2 == lado3 && lado1 != lado2 {
			fmt.Print("O triângulo é isóceles")
		}
	} else {
		fmt.Print("Os valores fornecidos não formam um triângulo")
	}
}
