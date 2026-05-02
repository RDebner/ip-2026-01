package main

import "fmt"

func main() {
	var num = []int{}
	var n int

	for len(num) < 10 {
		fmt.Scan(&n)
		num = append(num, n)
	}

	var pares, impares = []int{}, []int{}

	for i:=0; i < len(num); i++ {
		if num[i]%2 == 0 {
			pares = append(pares, num[i])
		} else {
			impares = append(impares, num[i])
		}
	}

	somaPar := 0 

	for i:=0; i < len(pares); i++ {
		somaPar += pares[i]
	}

	fmt.Printf("Números pares digitados: %d\nSoma dos números pares digitados: %d\nNúmeros ímpares digitados: %d\nQuantidade de números ímpares digitados: %d\n", pares, somaPar, impares, len(impares))
}