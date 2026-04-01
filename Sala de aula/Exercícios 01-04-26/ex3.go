package main

import "fmt"

var nume1, nume2, nume3 float64

func media (n1, n2, n3 float64) float64{
	return (n1 + n2 + n3) / 3
}

func main() {
	fmt.Print("Informe os valores: ")
	fmt.Scan(&nume1, &nume2, &nume3)
	fmt.Println("Média:", media(nume1,nume2,nume3))
}