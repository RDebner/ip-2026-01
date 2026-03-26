package main

import f "fmt"

var num int

func main() {
	f.Print("Escreva um número inteiro: ")
	f.Scan(&num)
	if num == 0 {
		f.Print("O número é nulo \n")
	} else 
	if num > 0 {
		f.Print("O número é positivo\n")
	} else
	if num < 0 {
		f.Print("O número é negativo\n")
	}
}