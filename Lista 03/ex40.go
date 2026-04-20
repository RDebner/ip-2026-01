package main

import (
	"fmt"
	"math"
)

var qntDescontos, lucroMax float64
type Lucros struct {
	preco float64
	lucro float64
}

var lucros Lucros
var listaLucros []Lucros

func main() {

	lucros.preco = 6
	qntDescontos = 0

	for {
		if lucros.preco < 1 {
			break
		}
		lucros.lucro = -18*math.Pow(qntDescontos,2) + 102*qntDescontos + 480
		listaLucros = append(listaLucros, lucros)
		fmt.Printf("Preço do ingresso: %.2f Lucro esperado: %.2f\n", lucros.preco, lucros.lucro)
		qntDescontos++
		lucros.preco = lucros.preco - 0.6
	} 
	
	indiceMaiorLucro := 0
	lucroMax = listaLucros[0].lucro

	for i:=0 ; i < len(listaLucros); i++{
		if listaLucros[i].lucro > lucroMax {
			lucroMax = listaLucros[i].lucro
			indiceMaiorLucro = i
		}
	}

	fmt.Printf("Maior lucro esperado: %.2f\nPreço: %.2f\nQuantidade de ingressos: %.2f", listaLucros[indiceMaiorLucro].lucro, listaLucros[indiceMaiorLucro].preco,(listaLucros[indiceMaiorLucro].lucro+300)/listaLucros[indiceMaiorLucro].preco)
}