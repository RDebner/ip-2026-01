package main

import "fmt"

var soma1ateP, p int

func main() {

	fmt.Print("Informe o valor de N: ")
	fmt.Scan(&p)

	for i:=1; i<=p; i++{
		soma1ateP += i
	}

	fmt.Print(soma1ateP)
}