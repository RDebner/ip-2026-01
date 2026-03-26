package main

import "fmt"

var destino, idaEVolta, valor int


func main() {
	fmt.Print("Informe o destino e o tipo de viagem: ")
	fmt.Scan(&destino, &idaEVolta)

	if destino == 0 || destino > 4 || idaEVolta < 0 || idaEVolta > 1 {
		fmt.Print("Valores inválidos")
		return
	}
	if destino == 1 {
		valor = 500 
		if idaEVolta == 1 {
			valor = 900
		}
	} else 
	if destino == 2 {
		valor = 350
		if idaEVolta == 1 {
			valor = 650
		}
	} else 
	if destino == 3 {
		valor = 350
		if idaEVolta == 1 {
			valor = 600
		}
	} else 
	if destino == 4 {
		valor = 300
		if idaEVolta == 1 {
			valor = 550
		}
	}

	fmt.Println("O valor da viagem é: ", valor)

}

