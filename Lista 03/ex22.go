package main

import "fmt"

var somaA int

func main() {
	j := 1
	
	for i:=38; i >= 2; i-- {
		somaA += (i*(i-1))/j
		j++
	}
	
	fmt.Println(somaA)
}