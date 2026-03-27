package main

import (
	f "fmt"
)

var numId int
var nota1, nota2, nota3, medEx, medAprov float64
var conc, aprovacao string

func main() {
	f.Print("Informe o número de identificação do aluno:")
	f.Scan(&numId)
	f.Print("Informe as 3 notas do aluno e a média dos exercícios:")
	f.Scan( &nota1, &nota2, &nota3, &medEx)

	medAprov = (nota1 + nota2*2 + nota3*3 + medEx) / 7

	if medAprov <= 10 && medAprov >= 9 {
		conc = "A"
	} else
	if medAprov < 9 && medAprov >= 7.5 {
		conc = "B"
	} else
	if medAprov < 7.5 && medAprov >= 6 {
		conc = "C"
	} else 
	if medAprov < 6 && medAprov >= 4 {
		conc = "D"
	} else {
		conc = "E"
	}

	switch conc {
	case "A", "B", "C":
		aprovacao = "APROVADO"
	case "D", "E":
		aprovacao = "REPROVADO"
	}

	f.Printf("Número do aluno: %d\nNotas: %.2f, %.2f, %.2f\nMédia dos exercícios: %.2f\nMédia de aproveitamento: %.2f\nConceito: %s - %s", numId, nota1, nota2, nota3, medEx, medAprov, conc, aprovacao)
}