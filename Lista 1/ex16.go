package main

import f "fmt"

var salar, salarReaj float64

func main() {
	f.Print("Informe o salário do funcionário: ")
	f.Scan(&salar)

	if salar <= 300 {
		salarReaj = salar*1.5
	} else {
		salarReaj = salar*1.3
	}

	f.Printf("SALARIO COM REAJUSTE = %.2f", salarReaj)
}