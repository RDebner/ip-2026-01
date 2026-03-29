package main

import "fmt"

var n1, n2, n3, num int

func main() {
	fmt.Print("Escreva 3 números inteiros com uma casa: ")
	fmt.Scan(&n1, &n2, &n3)

	if n1 > 9 || n1 == 0 || n2 > 9 || n3 > 9 {
		fmt.Print("DÍGITO INVÁLIDO")
		return
	}

	num = n1*100 + n2*10 + n3

	fmt.Printf("%d, %d", num, num*num)
}