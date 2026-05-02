package main

import "fmt"

var lis = [10]int{1,2,3,4,5,6,7,8,9,10}
var codigo int

func main() {
	fmt.Print("Informe o código: ")
	fmt.Scan(&codigo)
	if codigo == 0 {
		return
	} 
	if codigo == 1 {
		i := 0
		for i < 10 {
			if i == 9 {
				fmt.Printf("%d", lis[i])
			} else {
				fmt.Printf("%d, ", lis[i])
			}
			i++
		}	
	} 
	if codigo == 2 {
		i := 9
		for i >=0 {
			if i == 0 {
				fmt.Printf("%d", lis[i])
			} else {
				fmt.Printf("%d, ", lis[i])
			}
			i--
		}
	}

	
}