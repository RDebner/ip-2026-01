package main

import (
	f "fmt"
	"math"
)


var altura, aresta, areaB, volum float64

func main() {
	f.Print("Escreva o valor da altura e da aresta da base da pirâmide: ")
	f.Scan(&altura, &aresta)

	areaB = ((3*aresta*aresta)*math.Sqrt(3))/2
	volum = areaB*altura/3

	f.Printf("O VOLUME DA PIRÂMIDE E = %.2f METROS CÚBICOS", volum)
}