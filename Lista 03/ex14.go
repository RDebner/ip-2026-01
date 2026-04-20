package main

import "fmt"

var n1, n2, a int
var numPrimos = []int{}


func main() {
	fmt.Print("Informe o N1 e o N2: ")
	fmt.Scan(&n1,&n2)
	
	for {
		if n1 == n2 {
			break
		}

		for i:=1; i <= n1; i++ {
			if n1%i == 0 {
				a++
			}
		}

		if a == 2 {
			numPrimos = append(numPrimos, n1)
		}

		n1++
		a = 0
	}

	fmt.Print(numPrimos)
}