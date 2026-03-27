package main

import f "fmt"

var idade int
var classeEleit string

func main() {
	f.Print("Informe sua idade: ")
	f.Scan(&idade)

	if idade < 0 {
		f.Print("Idade inválida")
		return
	}

	if idade < 16 {
		classeEleit = "Não-eleitor"
	} else 
	if idade >= 16 && idade < 18 || idade >= 65{
		classeEleit = "Eleitor facultativo"
	} else
	if idade >= 18 && idade < 65 {
		classeEleit = "Eleitor obrigatório"
	}

	f.Printf("Classe eleitoral: %s", classeEleit)
}