package main

import "fmt"

var vetor = [30]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
var vetor2 [30]int

func main() {
	for i:=0 ; i < 30 ; i++ {
		if i%2 == 0 {
			vetor2[i] = vetor[i]*2
		} else {
			vetor2[i] = vetor[i]*3
		}
	}

	fmt.Print(vetor2)
}