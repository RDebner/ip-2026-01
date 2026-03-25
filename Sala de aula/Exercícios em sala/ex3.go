package main

import "fmt"

var qntAlunos int
var soma float64 = 0
var nota float64 

func main() {
	fmt.Print("Quantidade de alunos: ")
	fmt.Scan(&qntAlunos)
	for i := 0; i < qntAlunos; i++ {
		fmt.Printf("Escreva a nota do aluno %d:\n", i + 1)
		fmt.Scan(&nota)
		soma += nota
	}
	fmt.Println(soma/float64(qntAlunos))
}
