package main

import "fmt"

var c, d, m float64
var sequencia = []float64{1000}

func main() {

	fmt.Print("Informe a quantidade de termos: ")
	fmt.Scan(&m)

	c = sequencia[0]
	d = c-3

	for i:=2; i <= int(m); i++ {
		if i%2 == 0 {
			c -= d/float64(i)
		} else
		if i%2 != 0 {
			c += d/float64(i)
		}
		sequencia = append(sequencia, c)
		d-=3
	} 
	fmt.Print(sequencia)
}