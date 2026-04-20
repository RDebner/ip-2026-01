package main

import "fmt"

var matriz = [10][10]int{}

func main() {
	for i:=0; i < 10; i++ {
		for j:=0; j < 10; j++ {
			if j < i {
				fmt.Printf("[%d][%d] ", i,j)
			}
		}
		fmt.Println()
	}
}