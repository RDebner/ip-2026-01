package main

import (
	"fmt"
	"slices"
)

func main() {
	var a = []int{7, 2, 4, 3, 3, 5, 5, 4, 4, 4}

	type Num struct {
		num    int
		qntRep int
	}

	n := Num{}
	listaRepetidos := []Num{}

	for len(a) > 0 {
		i := 0
		numero := a[i]
		repeticao := 1
		for j := 1; j < len(a); j++ {
			if numero == a[j] {
				repeticao++
			}
			for k:=0; k < len(listaRepetidos);k++ {
				if numero == listaRepetidos[k].num {
					repeticao = 1
				}
			}
		}
		if repeticao > 1 {
			n.num = numero
			n.qntRep = repeticao
			listaRepetidos = append(listaRepetidos, n)
		}
		a = slices.Delete(a, 0, 1)
	}

	for i:=0; i < len(listaRepetidos); i++ {
		fmt.Printf("Número: %d  Quantidade de repetições: %d\n", listaRepetidos[i].num, listaRepetidos[i].qntRep)
	}

}
