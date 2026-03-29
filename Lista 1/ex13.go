package main

import f "fmt"

var nota float64
var conc string

func main() {
	f.Print("Escreva a nota: ")
	f.Scan(&nota)

	if nota >= 9 && nota <= 10 {
		conc = "A"
	} else
	if nota >= 7.5 && nota < 9 {
		conc = "B"
	} else
	if nota >= 6 && nota < 7.5 {
		conc = "C"
	} else {
		conc = "D"
	}
		
	f.Printf("NOTA = %.2f CONCEITO = %s", nota, conc)

}