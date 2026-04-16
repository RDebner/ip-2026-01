package main

import "fmt"

var nBase10 , rest int
var hex string
var lisRestos = []int{}

func main() {
	fmt.Print("Informe um número inteiro positivo na base 10: ")
	fmt.Scan(&nBase10)

	for nBase10 > 0 {
		rest = nBase10%16
		lisRestos = append(lisRestos, rest)
		nBase10 = nBase10 / 16
	}

	for i:=len(lisRestos)-1 ; i >=0 ; i-- {
		if lisRestos[i] >= 0 && lisRestos[i] <=9 {
			hex += fmt.Sprint(lisRestos[i])
		} else 
		if lisRestos[i] >= 10 && lisRestos[i] <= 15 {
			switch lisRestos[i] {
			case 10: hex += fmt.Sprint("A")
			case 11: hex += fmt.Sprint("B")
			case 12: hex += fmt.Sprint("C")
			case 13: hex += fmt.Sprint("D")
			case 14: hex += fmt.Sprint("E")
			case 15: hex += fmt.Sprint("F")
			}

		}
	}

	fmt.Printf("Hexadecimal: %s\n", hex)
}