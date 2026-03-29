package main

import f "fmt"

var numPar, qntNum int

var numPares []int

func main() {
	f.Print("Escreva o número par inicial e quantos deverão ser mostrados: ")
	f.Scan(&numPar, &qntNum)

	if numPar%2 != 0 {
		f.Print("O PRIMEIRO NÚMERO NÃO É PAR")
		return
	}

	for i := 0 ; i < qntNum ; i++ {
		numPares = append(numPares, numPar)
		numPar = numPar + 2
	}

	f.Println(numPares)
}