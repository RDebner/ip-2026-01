package main

import "fmt"

var qntNotas int
var soma float64 = 0
var nota float64 

func main() {
	fmt.Print("Quantidade de notas: ")
	fmt.Scan(&qntNotas)
	for i := 0; i < qntNotas; i++ {
		fmt.Printf("Escreva a nota %d:\n", i + 1)
		fmt.Scan(&nota)
		soma += nota
	}
	fmt.Println(soma/float64(qntNotas))
}