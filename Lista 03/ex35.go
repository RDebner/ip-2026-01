package main

import "fmt"

var numBase10, resto int
var binario string
var listaRestos = []int{}

func main() {
	fmt.Print("Informe um número inteiro positivo na base 10: ")
	fmt.Scan(&numBase10)
	
	for numBase10 > 0 {
		resto = numBase10%2
		listaRestos = append(listaRestos, resto)
		numBase10 = numBase10/2
	}

	for i:=len(listaRestos)-1 ; i >=0; i-- {
		binario += fmt.Sprint(listaRestos[i])
	}

	fmt.Printf("Binário: %s\n", binario)
}