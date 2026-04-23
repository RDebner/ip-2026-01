package main

import (
	"fmt"
	"strconv"
)

var restos = []int{}

func binario(n int){
	resto := n%2
	restos = append(restos, resto)
	if n == 1 {
		return
	}
	binario(n/2)
}

func main() {
	binario(10)
	var binario string
	for i:=len(restos)-1; i >=0 ; i-- {
		s := strconv.Itoa(restos[i]) 
		binario += s
	}
	fmt.Println(binario)
}