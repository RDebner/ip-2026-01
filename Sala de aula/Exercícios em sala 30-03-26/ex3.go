package main

import "fmt"

const num int = 10
var numer[num]int

func main() {
	for i:=0; i < num; i++ {
		fmt.Printf("Informe o número %d: ",i + 1)
		fmt.Scan(&numer[i])
	}

	for i := num; i > 0; i = i - 1 {
		fmt.Printf("%v ",numer[i - 1])
	}
	fmt.Printf("\n")
}