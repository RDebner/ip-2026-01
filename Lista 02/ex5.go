package main

import f "fmt"

var num int

func main() {
	f.Print("Escreva um número: ")
	f.Scan(&num)

	if num % 5 == 0 {
		f.Printf("O número %d é divisível por 5 \n", num)
	} else {
		f.Printf("O número %d não é divisível por 5 \n", num)
	}
}