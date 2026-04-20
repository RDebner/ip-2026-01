package main

import "fmt"

type Bois struct {
	numeroBoi int
	peso float64
}

var listaBois []Bois
var bois Bois
var maiorPeso, menorPeso float64
var boiMaisPesado, boiMaisLeve int

func main() {

	for i:=0; i<90; i++ {
		fmt.Printf("Informe o peso do boi número %d: \n", i+1)
		bois.numeroBoi = i+1
		fmt.Scan(&bois.peso)
		listaBois = append(listaBois, bois)
	}

	menorPeso = listaBois[0].peso

	for i:=0; i < len(listaBois); i++ {
		if listaBois[i].peso > maiorPeso {
			boiMaisPesado = listaBois[i].numeroBoi
			maiorPeso = listaBois[i].peso
		}
		if listaBois[i].peso < menorPeso {
			boiMaisLeve = listaBois[i].numeroBoi
			menorPeso = listaBois[i].peso
		}
	}

	fmt.Printf("Boi mais pesado: Número %d com peso %.2f\nBoi mais leve: Número %d com peso %.2f", boiMaisPesado, maiorPeso, boiMaisLeve, menorPeso)
}