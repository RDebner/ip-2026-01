package main

import (
	"fmt"
	"math"
)

var ss, xx, piatorio float64

func main() {
	xx = 1
	for i:=1; i <=51; i++ {
		if i%2 == 0 {
			ss -= 1/math.Pow(xx,3)
		} else 
		if i%2 != 0 {
			ss += 1/math.Pow(xx,3)
		}
		xx += 2
	}

	piatorio = math.Cbrt(ss*32)

	fmt.Print(piatorio)
}