package main

import "fmt"

var valorCompra, valorVenda float64

func main() {
	fmt.Print("Valor de compra: ")
	fmt.Scan(&valorCompra)
	
	if valorCompra == 0 {
		fmt.Print("Valor inválido \n")
		return
	}
	if valorCompra < 10 {
		valorVenda = valorCompra*1.7
	}
	if valorCompra >= 10 && valorCompra < 30{
		valorVenda = valorCompra*1.5
	}
	if valorCompra >= 30 && valorCompra < 50 {
		valorVenda = valorCompra*1.4
	}
	if valorCompra >= 50 {
		valorVenda = valorCompra*1.3
	}

	fmt.Println("Valor de venda:",valorVenda)

}

