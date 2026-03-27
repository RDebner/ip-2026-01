package main

import f "fmt"

var diaSem, categ string
var precoNorm, precoFin float64

func main() {
	f.Print("Informe o dia da semana: ")
	f.Scan(&diaSem)
	f.Print("Informe o preco e a categoria do DVD: ")
	f.Scan(&precoNorm, &categ)
	switch diaSem {
	case "Segunda", "Terça", "Quinta":
		precoFin = precoNorm*0.6
		if categ == "Lançamento" {
			precoFin = precoFin*1.15
		}
	case "Quarta", "Sexta", "Sábado", "Domingo":
		precoFin = precoNorm
		if categ == "Lançamento" {
			precoFin = precoFin*1.15
		}
	}

	f.Printf("O valor do aluguel é: %.2f", precoFin)
}