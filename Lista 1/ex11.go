package main

import f "fmt"

var num int

func main() {
	f.Print("Escreva um número: ")
	f.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		f.Print("O NÚMERO E DIVISÍVEL")
	} else {
		f.Print("O NÚMERO NÃO É DIVISÍVEL")
	}
}