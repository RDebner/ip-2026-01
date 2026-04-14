package main

import "fmt"

var x, n int

func potencia(x, n int) int {
	if n == 1 {
		return x
	}

	return x * potencia(x, n-1)
}

func main() {
	fmt.Print("Informe a base e o expoente: ")
	fmt.Scan(&x,&n)

	fmt.Println(potencia(x,n))
}