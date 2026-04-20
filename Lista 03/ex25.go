package main

import "fmt"

var t, v, somaT float64

func main() {

	t, v = 1,15

	for i:=1; i <= 15; i++ {
		if i%2 != 0 {
			somaT += t/(v*v)
		} else
		if i%2 == 0 {
			somaT -= t/(v*v)
		}
		t *= 2
		v--
		fmt.Println(somaT)
	}
}