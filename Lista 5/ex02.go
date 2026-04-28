package main

import "fmt"

func main() {
	var v1, v2 = []int{}, []int{}
	var n int
	
	fmt.Print("Informe os valores do primeiro vetor: ")
	for len(v1) < 10 {
		fmt.Scan(&n)
		v1 = append(v1, n)
	}
	fmt.Print("Informe os valores do segundo vetor: ")
	for len(v2) < 5 {
		fmt.Scan(&n)
		v2 = append(v2, n)
	}

	var par, impar = []int{}, []int{} 

	for i:= 0; i < len(v1); i++ {
		if v1[i]%2 == 0 {
			par = append(par, v1[i])
		} else {
			impar = append(impar, v1[i])
		}
	}

	var vr1, vr2 = []int{}, []int{}

	for i:= 0; i < len(par) ; i++ {
		soma := par[i]
	  	for j := 0 ; j < len(v2); j++ {
			soma += v2[j]
		}
		vr1 = append(vr1, soma)		
	}

	for i:= 0; i < len(impar) ; i++ {
		soma := impar[i]
	  	for j := 0 ; j < len(v2); j++ {
			soma += v2[j]
		}
		vr2 = append(vr2, soma)		
	}

	fmt.Println(vr1, vr2)
}