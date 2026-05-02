package main

import (
	"fmt"
	f "fmt"
)

func main() {
	var alturas = [10]float64{1.59, 1.6, 1.72, 1.68, 1.66, 1.8, 1.83, 1.9, 1.82, 1.8}
	var somaAlt float64

	for i:=0; i < len(alturas); i++ {
		somaAlt += alturas[i]
	}

	media := somaAlt / float64(len(alturas))

	var alturaAcimMedia = []float64{}

	for i:=0; i < len(alturas); i++ {
		if alturas[i] > media {
			alturaAcimMedia = append(alturaAcimMedia, alturas[i])
		}
	}
	
	imprimir(alturaAcimMedia)
}

func imprimir(arr []float64)  {
	for i, value := range arr {
		if i == len(arr)-1 {
			fmt.Printf("%.2f", value)
		} else {
			f.Printf("%.2f, ", value)
		}
	}
}