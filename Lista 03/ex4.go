package main

import "fmt"

var num float64

func main() {

	for{
		fmt.Print("Informe um número: ")
		fmt.Scan(&num)

		if num <= 0 {
			return
		}

		for i:= 1; i <= int(num); i++{
			if num/float64(i) == float64(i) {
				fmt.Printf("O número %d é um quadrado perfeito\n", int(num))
				break
			}
			if i == int(num) {
				fmt.Printf("O número %d não é um quadrado perfeito\n-/", int(num))
				break
			}
		}
	}
}