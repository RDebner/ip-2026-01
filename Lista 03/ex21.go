package main

import "fmt"

var baseb, expoente, potencia int

func main() {
	fmt.Print("Informe a base b e o expoente n (b >= 2, n > 1): ")
	fmt.Scan(&baseb, &expoente)

	potencia = baseb

	for i:=1; i < expoente; i++ {
		potencia *= baseb
	}

	fmt.Printf("%d", potencia)
}