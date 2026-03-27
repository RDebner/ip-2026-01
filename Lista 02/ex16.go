package main

import (
	f "fmt"
	"math"
)

var a, b, c float64 
var raizes string

func main() {
	f.Print("Escreva os coeficientes da função: ")
	f.Scan(&a, &b, &c)

	var delta float64 = (b*b) - (4*a*c)
	var raiz1, raiz2 float64
	
	if delta == 0 {
		raiz1 = -(b) / 2*a
		raizes = "RAIZ ÚNICA"
	} else
	if delta > 0 {
		raiz1 = (-b + math.Sqrt(delta)) / 2*a
		raiz2 = (-b - math.Sqrt(delta)) / 2*a
		raizes = "RAÍZES DISTINTAS"
	} else
	if delta < 0 {
		raizes = "RAÍZES IMAGINÁRIAS"
		f.Println(raizes)
		return
	}

	f.Println(raiz1, raiz2, raizes)
	
}