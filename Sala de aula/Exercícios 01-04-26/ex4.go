package main

import "fmt"

var number int

func fatorial(n int) int {
	var fatorial int = n
	for i := n - 1; i > 0; i-- {
		fatorial = fatorial * i
	}
	return fatorial
}

func main() {
	fmt.Print("Informe um número:")
	fmt.Scan(&number)
	fmt.Println(fatorial(number))
}