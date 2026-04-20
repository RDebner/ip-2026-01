package main

import "fmt"

var n int
var fibonacci = []int{0,1}

func main() {
	fmt.Print("Informe o número de termos N da sequência de Fibonacci (N >= 3): ")
	fmt.Scan(&n)

	for i:=2; i < n; i++ {
		var proxTermo = fibonacci[i-1] + fibonacci[i-2]
		fibonacci = append(fibonacci, proxTermo)
	}

	for i, v:= range fibonacci {
		fmt.Print(v)
		if i < len(fibonacci)-1 {
			fmt.Print(", ")
		}
	}
}