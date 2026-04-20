package main

import "fmt"

var numer, somaNum, maior, menor, somaPar, mediaPar, porcImpar, mediaNumeros float64
var listaNumeros = []float64{}
var par = []float64{}
var impar = []float64{}

func main() {

	fmt.Print("Digite os números (digite 30.000 para finalizar): ")

	for {
		fmt.Scan(&numer)
		if numer == 30000 {
			break
		}
		listaNumeros = append(listaNumeros, numer)
	}


	for i:=0; i < len(listaNumeros); i++ {

		somaNum += listaNumeros[i]

		menor = listaNumeros[0]

		if listaNumeros[i] >= maior {
			maior = listaNumeros[i]
		}
		if listaNumeros[i] <= menor {
			menor = listaNumeros[i]
		}

		if int(listaNumeros[i])%2 == 0 {
			par = append(par, listaNumeros[i])
		} else {
			impar = append(impar, listaNumeros[i])
		}
	}

	for i:= 0; i < len(par); i++ {
		somaPar += par[i]
	}



	mediaNumeros = somaNum/float64(len(listaNumeros))
	mediaPar = somaPar/float64(len(par))
	porcImpar = float64(len(impar))/float64(len(listaNumeros))*100

	fmt.Printf("Soma dos números digitados: %.2f\nQuantidade de números digitados: %.0f\nMédia dos números digitados: %.2f\nMaior número digitado: %.0f\nMenor número digitado: %.0f\nMédia dos números pares: %.2f\nPorcentagem de números ímpares: %.0f\n", somaNum, float64(len(listaNumeros)), mediaNumeros, maior, menor, mediaPar, porcImpar)
}