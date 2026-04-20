package main

import "fmt"

var qntTermos int

func main() {
	fmt.Print("Informe a quantidade de termos: ")
	fmt.Scan(&qntTermos)

	for i:=1; i <= qntTermos; i++ {
		if i < qntTermos {
			fmt.Printf("%d, ", i*i)
		}
		if i == qntTermos {
			fmt.Printf("%d", i*i)
		}
	}
}