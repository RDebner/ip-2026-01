package main

import "fmt"

var soma int

func main() {
	for i:=1; i <= 20; i++ {
		soma += i
		fmt.Println(i)
	}
	fmt.Print(soma)
}