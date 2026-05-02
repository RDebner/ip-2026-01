package main

import "fmt"

var listaIdades = [50]int{23, 45, 12, 67, 34, 55, 8, 29, 48, 33, 45, 19, 62, 48, 62, 62, 62, 62, 50, 48, 66, 22, 44, 30, 11, 58, 72, 62, 19, 47, 60, 25, 39, 14, 52, 70, 31, 28, 43, 6, 56, 17, 35, 69, 21, 48, 75, 10, 32, 59}

type Idade struct {
	idade int
	rep int
}

var idd = Idade{}
var idades = []Idade{}

func main() {
	idd.idade = listaIdades[0]
	idades = append(idades, idd)
	for i:=1; i < len(listaIdades); i++ {
		acres := true
		for j:=0; j < len(idades); j++ {
			if listaIdades[i] == idades[j].idade {
				acres = false
			}
		}
		if acres == true {
			idd.idade = listaIdades[i]
			idades = append(idades, idd)
		}
	}

	for i:=0; i < len(idades); i++ {
		for j:=0; j < len(listaIdades); j++ {
			if idades[i].idade == listaIdades[j] {
				idades[i].rep += 1
			}
		}
	}

	moda := [1]Idade{}
	moda[0] = idades[0]

	for i:=1; i < len(idades); i++ {
		if idades[i].rep > moda[0].rep {
			moda[0] = idades[i]
		}
	}

	fmt.Printf("A moda das idades é %d anos", moda[0].idade)

}