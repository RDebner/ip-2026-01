package main

import "fmt"

var numeros[5]int
var soma int

func main() {
	fmt.Print("Escreva os 5 números a serem somados: ")
	fmt.Scan(&numeros[0], &numeros[1], &numeros[2], &numeros[3], &numeros[4])

	for i := 0; i < len(numeros); i++ {
		soma += numeros[i]
	}

	fmt.Println(soma)

}
