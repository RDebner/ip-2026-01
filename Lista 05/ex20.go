package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	sortear := rand.New(rand.NewSource(time.Now().UnixNano()))

	var numerosSorteados = [20]int{}

	for i:=0; i < len(numerosSorteados); i++ {
		numSort := sortear.Intn(6) + 1
		numerosSorteados[i] = numSort
	}

	type Sorteados struct {
		num int
		repet int
	}

	var sorteados = [6]Sorteados{{1,0},{2,0},{3,0},{4,0},{5,0},{6,0}}

	for i:=0; i < len(sorteados); i++ {
		for j:=0; j < len(numerosSorteados); j++ {
			if sorteados[i].num == numerosSorteados[j] {
				sorteados[i].repet++
			}
		}
	}

	fmt.Printf("%d\n", numerosSorteados)

	for i:=0; i < len(sorteados); i++ {
		fmt.Printf("Frequência do número %d: %d\n", sorteados[i].num, sorteados[i].repet)
	}
}