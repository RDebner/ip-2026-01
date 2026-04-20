package main

import (
	"fmt"
	"math"
)

var u float64

func serieMacLaurin(a float64) float64 {
	return a - (math.Pow(a,3)/6) + (math.Pow(a,5)/120) - (math.Pow(a,7)/5040)
} 

func main() {
	for u <= 6.3 {
		fmt.Println(serieMacLaurin(u))
		u = u + 0.1
	}
}