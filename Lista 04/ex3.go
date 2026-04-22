package main

import "fmt"

var inteiros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func inverter(l []int) {
	i := len(l)-1
	if i <= 0 {
		return
	}
	l[i], l[0] = l[0], l[i]
    inverter(l[1:i])
}
	

func main() {
	inverter(inteiros)
	fmt.Println(inteiros)
}
