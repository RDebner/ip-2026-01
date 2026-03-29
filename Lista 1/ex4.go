package main

import "fmt"

var salmin, qntKw, custoCons float64

func main() {
	fmt.Print("Informe o salário mínimo: ")
	fmt.Scan(&salmin)
	fmt.Print("Informe a quantidade de kW gasta: ")
	fmt.Scan(&qntKw)

	custoCons = (0.7*salmin*qntKw)/100

	fmt.Printf("Custo por kW: R$%.2f\nCusto do consumo: R$%.2f\nCusto com desconto: R$%.2f", custoCons/qntKw, custoCons, custoCons*0.9)
}