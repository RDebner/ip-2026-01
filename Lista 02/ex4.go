package main

import f "fmt"
import m "math"

var num float64

func main() {
	f.Print("Escreva um número: ")
	f.Scan(&num)
	if num >= 0 {
		f.Println(m.Sqrt(num))
	} else {
		f.Println(num*num)
	}
}