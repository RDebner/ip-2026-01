package main

import "fmt"

var soma, total, media int

func main() {
	for i := 50; i <= 70; i = i + 2 {
		soma += i
		total++
	}
	media = soma / total

	fmt.Printf("Soma = %d\nMédia = %d", soma, media)
}