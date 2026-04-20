package main

import "fmt"

var numInt int

func main() {
	fmt.Print("Informe um número inteiro: ")
	fmt.Scan(&numInt)
	
	for i:=1 ; i + 2 <= numInt; i++ {
		if i*(i+1)*(i+2) == numInt {
			fmt.Printf("O número %d é um número triangular.", numInt)
			return
		}
		if  i + 2 == numInt {
			fmt.Printf("O número %d não é um número triangular.", numInt)
			return
		}
	} 
}