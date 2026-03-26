package main

import "fmt"

var num int

func main() {
	fmt.Print("Escreva um número: ")
	fmt.Scan(&num)
	if num - 90 > -70 && num - 90 < 0 {
		fmt.Printf("O número %d está compreendido entre 20 e 90\n", num)
	} else {
		fmt.Printf("O número %d não está compreendido entre 20 e 90\n", num)
	}
}