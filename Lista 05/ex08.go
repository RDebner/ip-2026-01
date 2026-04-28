package main

import "fmt"

func main() {
	var quadrado = [15]int{}
	var num int

	for i:=0; i < 15; i++ {
		fmt.Scan(&num)
		if num >= 0 {
			quadrado[i] = num * num
		} else {
			quadrado[i] = -1
		}
	}

	fmt.Print(quadrado)
}