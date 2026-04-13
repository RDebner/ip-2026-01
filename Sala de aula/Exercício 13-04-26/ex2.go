package main

import "fmt"

func buscaBinaria(l []int, x int) int {
	var n, e, d int
	n = len(l)
	e = 0
	d = n - 1
	for e <= d {
		m := (e + d)/2
		if l[m] == x {
			return  m
		} 
		if l[m] < x {
			e = m + 1
		} 
		if l[m] > x {
			d = m - 1
		}
	}
	return -1
}

func main() {
	var lista2 = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,4,15}
	fmt.Println(buscaBinaria(lista2, 5))
}