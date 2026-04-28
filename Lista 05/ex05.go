package main

import "fmt"

func main() {

	var vetor = [10]int{1,2,3,4,5,6,7,8,9,0}
	var menor, indice int

	menor = vetor[0]
	indice = 0

	for i:=0; i < len(vetor); i++ {
		if vetor[i] < menor {
			menor = vetor[i]
			indice = i
		}
	}

	fmt.Printf("O menor elemento do vetor é %d e sua posição dentro do vetor é %d", menor, indice)
}