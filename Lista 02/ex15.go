package main

import f "fmt"

var dia, mes, ano int
var nomeMes string

func main() {
	f.Print("Informe a data, o mês e o ano: ")
	f.Scan(&dia, &mes, &ano)

	if dia == 0 || dia < 0 || dia > 31 {
		f.Print("Data inválida")
		return
	}
	switch mes {
	case 1:
		nomeMes = "Janeiro"
	case 2:
		nomeMes = "Fevereiro"
	case 3:
		nomeMes = "Março"
	case 4:
		nomeMes = "Abril"
	case 5:
		nomeMes = "Maio"
	case 6:
		nomeMes = "Junho"
	case 7:
		nomeMes = "Julho"
	case 8:
		nomeMes = "Agosto"
	case 9:
		nomeMes = "Setembro"
	case 10:
		nomeMes = "Outubro"
	case 11:
		nomeMes = "Novembro"
	case 12:
		nomeMes = "Dezembro"
	}
	f.Printf("%d de %s de %d", dia, nomeMes, ano)
}
