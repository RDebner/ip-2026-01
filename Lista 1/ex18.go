package main

import f "fmt"

var valInic, razao, qntTerm, soma int
var termos []int

func main() {
	f.Print("Escreva o valor inicial, a razão e a quantidade de termos da PA: ")
	f.Scan(&valInic, &razao, &qntTerm)

	termos = append(termos, valInic)
	
	for i:= 1 ; i < qntTerm ; i++ {
		termos = append(termos, termos[i - 1] + razao)
	}
	
	for i := 1; i <= qntTerm ; i++ {
		soma = soma + termos[i - 1]
	}

	f.Print(soma)
}