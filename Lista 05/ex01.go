package main

import "fmt"

func main() {
	var inteiros = []int{}
	var num int

	for len(inteiros) < 10 {
		fmt.Scan(&num)
		inteiros = append(inteiros, num)
	}

	type Maior50 struct {
		numero int
		indice int
	}

	maior50 := Maior50{}
	listaMaior50 := []Maior50{}

	for i:=0 ; i < 10; i++ {
		if inteiros[i] > 50 {
			maior50.numero = inteiros[i]
			maior50.indice = i
			listaMaior50 = append(listaMaior50, maior50)
		}
	}

	for i:=0 ; i < len(listaMaior50); i++ {
		fmt.Printf("Número: %d  Índice: %d \n", listaMaior50[i].numero, listaMaior50[i].indice)
	}

	if len(listaMaior50) == 0 {
		fmt.Print("Não tem nenhum número maior que 50.")
	}
}
