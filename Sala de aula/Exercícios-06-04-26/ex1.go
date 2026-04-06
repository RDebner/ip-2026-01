package main

import "fmt"


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
	
	var pessoa Pessoa

	for pessoa.nome != "FIM" {
		fmt.Printf("Escreva o nome da pessoa: ")
		fmt.Scan(&pessoa.nome)
		if pessoa.nome == "FIM" {
			for i:=0; i < len(listaPessoas); i++{
				fmt.Printf("O peso ideal de %s é %.2fkg\n",listaPessoas[i].nome, listaPessoas[i].pesoIdeal)
			}
			return
		}
		fmt.Printf("Escreva a altura da pessoa: ")
		fmt.Scan(&pessoa.altura)
		pessoa.pesoIdeal = pesoIdeal(pessoa.altura)
		listaPessoas = append(listaPessoas, pessoa)
	}
	
	
}