package main

import f "fmt"

var matFunc int
var qntHrsEx, salarBrut float64

var salarMin float64 = 788
var valorHrEx float64 = 10

func main() {
	f.Print("Matrícula do funcionário: ")
	f.Scan(&matFunc)
	f.Print("Quantidade de horas-extras trabalhadas: ")
	f.Scan(&qntHrsEx)

	var salarHrEx float64 = qntHrsEx*valorHrEx
	salarBrut = 3*salarMin + salarHrEx
	var salarLiq = salarBrut - (salarBrut*0.12) - (salarBrut*0.2)

	f.Printf("Salário Bruto: %.2f\nSalário Líquido: %.2f", salarBrut, salarLiq)
}