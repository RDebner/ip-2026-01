package main

import "fmt"

var salCarlos, salJoao, totalMeses float64

func main() {
	fmt.Print("Informe o salário de Carlos: ")
	fmt.Scan(&salCarlos)
	salJoao = salCarlos/3

	for {
		if salJoao >= salCarlos {
			break
		}
		totalMeses++
		salCarlos *= 1.02
		salJoao *= 1.05
	}

	fmt.Printf("João precisa de %.2f meses para igualar ou ultrapassar o valor de Carlos", totalMeses)
}