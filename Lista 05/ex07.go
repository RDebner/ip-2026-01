package main

import "fmt"

func main() {
	var impares = [100]int{}
	x := 1

	for i:=0; i < 100; i++ {
		impares[i] = x
		x +=2
	}

	fmt.Print(impares)

}