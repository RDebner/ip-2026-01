package main

import "fmt"

var numeros = [10]int{1,2,3,4,5,6,7,8,9,10}

type Primos struct {
	num int
	ind int
}

var p = Primos{}
var listaPrimos = []Primos{}

func main() {
	for i:=0; i < len(numeros); i++ {
		contagem := 0
		for j:=1; j <= numeros[i]; j++ {
			if numeros[i]%j == 0 {
				contagem++
			}
		}
		if contagem == 2 {
			p.num = numeros[i]
			p.ind = i
			listaPrimos = append(listaPrimos, p)
		}
	}
	
	for i:=0; i < len(listaPrimos); i++ {
		fmt.Printf("N. Primo: %d  Posição: %d\n", listaPrimos[i].num, listaPrimos[i].ind)
	}
}