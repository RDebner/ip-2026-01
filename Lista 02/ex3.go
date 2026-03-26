package main

import f "fmt"

var num1, num2 int

func main() {	
	f.Print("Escreva o primeiro número e o segundo número: ")
	f.Scan(&num1, &num2)
	var soma int = num1 + num2
	if soma > 20 {
		soma = soma + 8
		f.Println(soma)
	} else 
	if soma <= 20 {
		soma = soma - 5
		f.Println(soma)
	}
}