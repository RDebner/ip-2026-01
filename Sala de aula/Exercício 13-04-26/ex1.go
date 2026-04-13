package main

import "fmt"

var lista = []int{4,5,6,7,8}
var num, indice int

func buscaSequencial(l []int, x int) int {
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			indice = i
		} else {
			indice = -1
		}
	}
	return indice
} 

func main() {
	fmt.Print("Informe o número a ser buscado: ")
	fmt.Scan(&num)
	fmt.Println(buscaSequencial(lista, num))

}