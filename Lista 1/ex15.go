package main

import f "fmt"

var num int

func main() {
	f.Print("Escreva um número entre 5 e 2000: ")
	f.Scan(&num)
	for i := 2; i <= num ; i = i + 2 {
		f.Printf("%d ^ %d = %d\n", i, i, i*i)
	}
}