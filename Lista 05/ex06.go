package main

import "fmt"

func main() {
	var b = [100]int{}
	num := 100

	for i:=0 ; i < 100; i++ {
		b[i] = num
		num--
	}

	fmt.Print(b)
}