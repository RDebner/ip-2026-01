package main

import (
	"fmt"
	"math"
)

var volume, r, pi float64

func main() {

	r, pi = 0, 3.14

	for r <= 20 {
		volume = (4*pi*math.Pow(r,3))/3
		fmt.Println(volume)
		r += 0.5
	}
}