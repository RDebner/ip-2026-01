package main

import "fmt"

var numPessoas int

type Pessoa struct {
	nome string
	altura float64
	pesoIdeal float64
}

var listaPessoas = []Pessoa{}

func pesoIdeal(alt float64) float64 {
	return (72.7*alt)-58
}

func main() {
	fmt.Print("Informe o número de pessoas: ")
	fmt.Scan(&numPessoas)
	for i:=0; i < numPessoas; i++ {
		var pessoa Pessoa
		fmt.Printf("Escreva o nome da pessoa %d: ", i + 1)
		fmt.Scan(&pessoa.nome)
		fmt.Printf("Escreva a altura da pessoa %d: ", i + 1)
		fmt.Scan(&pessoa.altura)
		pessoa.pesoIdeal = pesoIdeal(pessoa.altura)
		listaPessoas = append(listaPessoas, pessoa)
	}
	for i:=0; i < numPessoas; i++{
		fmt.Printf("O peso ideal de %s é %.2f\n",listaPessoas[i].nome, listaPessoas[i].pesoIdeal)
	}
}