package main

import "fmt"

var x, y int

func potencia(x, y int) int {
	if  y == 1{
		return x
	}
	return x * potencia(x , y-1)
}

func main() {
	fmt.Print("Informe a base e o expoente: ")
	fmt.Scan(&x, &y)
	fmt.Print(potencia(x , y))
}