package main

import f "fmt"

var idade int


func main() {
	f.Print("Escreva sua idade: ")
	f.Scan(&idade)
	if idade > 0 && idade <= 2 {
		f.Print("Recém-nascido")
	}
	if idade >= 3 && idade <= 11 {
		f.Print("Criança")
	}
	if idade >= 12 && idade <= 19 {
		f.Print("Adolescente")
	}
	if idade >= 20 && idade <= 55 {
		f.Print("Adulto")
	}
	if idade > 55 {
		f.Print("Idoso")
	}
}
