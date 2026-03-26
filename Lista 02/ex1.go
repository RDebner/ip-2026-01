package main

import f "fmt" 

var num int

func main() {
	f.Print("Escreva um número inteiro: ")
	f.Scan(&num)

	if num % 2 == 0 {
		f.Print("O número e par \n")
	} else {
		f.Print("O número é impar \n")
	}
}