package main

import "fmt"

var lisNum = [10]int{5, 12, 4, 7, 10, 3, 2, 6, 23, 16}
var div = [5]int{3, 11, 5, 8, 2}

type Num struct {
	num int
	divisores []int
	indices []int
}

var listaNumber = []Num{}

func main() {
	for i := 0; i < 10; i++ {
		number := new(Num)
		number.num = lisNum[i]
		for j:=0; j < 5; j++ {
			if lisNum[i]%div[j] == 0 {
				number.divisores = append(number.divisores, div[j])
				number.indices = append(number.indices, j)
			}
		}
		listaNumber = append(listaNumber, *number)
	}

	for i:=0; i < len(listaNumber); i++ {
		fmt.Printf("Número %d: \n", listaNumber[i].num)
		if len(listaNumber[i].divisores) == 0 {
			fmt.Print("\tNão divisível por nenhum número da lista\n")
		} else {
			for j:=0; j < len(listaNumber[i].divisores); j++ {
				fmt.Printf("\tDivisível por %d na posição %d\n", listaNumber[i].divisores[j], listaNumber[i].indices[j])
			}
		}

	}
}