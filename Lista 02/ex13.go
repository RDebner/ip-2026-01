package main

import "fmt"

var num int 

func main() {
	fmt.Print("Escreva um número positivo de 3 casas: ")
	fmt.Scan(&num)
	if num < 0 {
		fmt.Print("O número informado não é positivo")
		return
	} 
	if num < 100 {
		fmt.Print("O número informado não possui 3 casas")
		return
	}

	var centenas int = num /100
	var dezena int = (num - centenas*100) / 10

	fmt.Println(dezena)

}