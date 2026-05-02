package main

import "fmt"

func main() {
	var vet1, vet2 [10]int

	fmt.Print("Informe os valores do vetor 1: ")
	for i:=0 ; i < 10 ; i++ {
		fmt.Scan(&vet1[i])
	}

	fmt.Print("Informe os valores do vetor 2: ")
	for i:=0 ; i < 10 ; i++ {
		fmt.Scan(&vet2[i])
	}

	var vetRes []int

	for i:= 0; i < 10; i++ {
		vetRes = append(vetRes, vet1[i])
		vetRes = append(vetRes, vet2[i])
	}

	fmt.Print(vetRes)
}