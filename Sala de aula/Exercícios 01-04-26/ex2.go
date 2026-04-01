package main

import "fmt"

var n1, n2, n3 int

func identificarMaiorValor(n1, n2, n3 int) int {
	var maiorNum int
	if n1 > n2 && n1 > n3 {
		maiorNum = n1
	} else 
	if n2 > n1 && n2 > n3 {
		maiorNum = n2
	} else 
	if n3 > n1 && n3 > n2 {
		maiorNum = n3
	}

	return maiorNum
}

func main() {
	fmt.Print("Informe 3 números: ")
	fmt.Scan(&n1, &n2, &n3)
	fmt.Println("Maior valor:", identificarMaiorValor(n1, n2, n3))
}