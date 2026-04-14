package main

import "fmt"

var qntTotalMilho, graos uint64

func main() {
	graos = 1
	for i := 0; i < 64; i++ {
		qntTotalMilho += graos
		graos *= 2
	}

	fmt.Println(qntTotalMilho)
}