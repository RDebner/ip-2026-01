package main

import "fmt"

const qntNotas int = 3
var notas[3] float64
var media float64
var notasAcimMedia[]float64
var notasAbaixMedia[]float64

func main() {
	fmt.Print("Informe as 3 notas do aluno: ")
	fmt.Scan(&notas[0], &notas[1], &notas[2])

	media = (notas[0] + notas[1] + notas[2]) / 3
	for i := 0; i < qntNotas; i++ {
		if notas[i] < media {
			notasAbaixMedia = append(notasAbaixMedia, notas[i])
		} else
		if notas[i] >= media {
			notasAcimMedia = append(notasAcimMedia, notas[i])
		}
	}

	fmt.Printf("Média: %.2f\nNotas acima da média: %v\nNotas abaixo da média: %v\n", media, notasAcimMedia, notasAbaixMedia)


}