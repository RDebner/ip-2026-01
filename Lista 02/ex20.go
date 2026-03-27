package main

import f "fmt"

var valorProd float64
var metPag string

func main() {
	f.Print("Informe o valor do produto:")
	f.Scan(&valorProd)
	f.Print("Informe a forma de pagamento:")
	f.Scan(&metPag)
	switch metPag {
	case "Dinheiro", "Cheque":
		valorProd = valorProd*0.9
		f.Printf("Preço: %.2f", valorProd)
	case "Crédito":
		valorProd = valorProd*0.95
		f.Printf("Preço: %.2f", valorProd)
	case "2 vezes":
		valorProd = valorProd / 2
		f.Printf("Duas parcelas de %.2f", valorProd)
	case "3 vezes":
		valorProd = valorProd*1.1
		f.Printf("Três parcelas de %.2f", valorProd)
}
}