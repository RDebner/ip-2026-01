package main

import "fmt"

var somaPrim20 float64

func calcFatorial(a float64) float64 {
	for i:=a - 1 ; i > 0; i-- {
		a *= i
	}
	if a == 0 {
		a = 1
	}
	return a
}
func main() {
	for i:=0 ; i<20; i++ {
		somaPrim20 += (100-float64(i))/calcFatorial(float64(i))
	}
	fmt.Println(somaPrim20)
}