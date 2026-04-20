package main

import "fmt"

var number, fatorial int


func main() {
	fmt.Print("Informe o número que deseja calcular o fatorial: ")
	fmt.Scan(&number)

	fatorial = number

	for i:=number-1; i > 0; i-- {
		fatorial *= i
	}

	fmt.Println(fatorial)
}