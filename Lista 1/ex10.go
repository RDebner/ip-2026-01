package main

import f "fmt"

var a, b, c, d, determinante float64

func main() {
	f.Print("Escreva os elementos a, b, c e d da matriz quadrada: ")
	f.Scan(&a, &b, &c, &d)

	determinante = a*d - b*c

	f.Printf("O VALOR DO DETERMINANTE E = %.2f", determinante)
}