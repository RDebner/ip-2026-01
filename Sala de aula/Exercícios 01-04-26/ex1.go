package main

import "fmt"

var num1, num2, somaNum int

func soma(n1, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Print("Informe dois valores: ")
	fmt.Scan(&num1, &num2)
	fmt.Println("Soma:",soma(num1, num2))
}