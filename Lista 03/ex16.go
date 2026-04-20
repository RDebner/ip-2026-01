package main

import "fmt"

var termo1, termo2, numTermos, proxTermo int

var listaFetuccine = []int{}

func main() {
	fmt.Print("Informe a quantidade N de termos (N >= 3): ")
	fmt.Scan(&numTermos)
	fmt.Print("Informe os dois primeiros termos: ")
	fmt.Scan(&termo1, &termo2)

	listaFetuccine = append(listaFetuccine, termo1, termo2)

	for i:=2; i < numTermos; i++ {
		if (i+1)%2 != 0 {
			proxTermo = listaFetuccine[i-1] + listaFetuccine[i-2]
		} else
		if (i+1)%2 == 0 {
			proxTermo = listaFetuccine[i-1] - listaFetuccine[i-2]
		}

		listaFetuccine = append(listaFetuccine, proxTermo)
	}

	fmt.Print(listaFetuccine)
}