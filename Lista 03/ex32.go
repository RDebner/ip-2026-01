package main

import "fmt"

var n1, n2, multiplicacao int

func main() {
	fmt.Print("Informe N1 E N2: ")
	fmt.Scan(&n1, &n2)


	for i:=0; i<n2; i++ {
		multiplicacao += n1
	}

	fmt.Println(multiplicacao)
}