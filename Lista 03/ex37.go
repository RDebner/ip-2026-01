package main

import (
	"fmt"
	"math"
	"strconv"
)

var numBase8, decimal int

func main() {
	fmt.Print("Informe um número inteiro positivo na base 8: ")
	fmt.Scan(&numBase8)

	n8 := strconv.Itoa(numBase8)
	
	numCasas := len(n8) -1

	for i:=0 ; i < len(n8); i++ {
		casaDecimal := numBase8/int(math.Pow(10,float64(numCasas)))
		decimal += casaDecimal*int(math.Pow(8,float64(numCasas)))
		numBase8 -= casaDecimal*int(math.Pow(10,float64(numCasas)))
		numCasas--
	}

	fmt.Print(decimal)
}