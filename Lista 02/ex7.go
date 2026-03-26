package main

import	"fmt"


var a, b, c, menor, inter, maior int


func main() {
	fmt.Print("Escreva 3 valores distintos: ")
	fmt.Scan(&a, &b, &c)

	if a < b && a < c {
		menor = a
		if b < c {
			inter = b
			maior = c
		} else {
			inter = c 
			maior = b
		}
	} else 
	if b < a && b < c {
		menor = b
		if a < c {
			inter = a 
			maior = c
		} else {
			inter = c
			maior = a
		}
	}
	if c < a && c < b {
		menor = c
		if a < b {
			inter = a
			maior = b
		} else { 
			inter = b
			maior = a
		}
	}

	fmt.Println(menor, inter, maior)
}