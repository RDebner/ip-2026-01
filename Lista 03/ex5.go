package main

import "fmt"

var idade, altura, peso, valor, mais50, menos40Kg, totalPessoas, mediaAlt, somaAlt, porcMenos40kg float64

var listaAltura = []float64{}

func main() {

	for {
		fmt.Print("Informe a idade, altura e peso da pessoa: ")
		fmt.Scan(&idade, &altura, &peso)

		totalPessoas++

		if idade > 50 {
			mais50++
		}
		if idade >= 10 && idade <= 20 {
			listaAltura = append(listaAltura, altura)
		}
		if peso < 40 {
			menos40Kg++
		}

		fmt.Print("Deseja inserir mais dados? Se sim digite 1, se não digite outro valor: ")
		fmt.Scan(&valor)
		if valor != 1 {
			break
		}
	}

	for i:= 0; i < len(listaAltura); i++ {
		somaAlt += listaAltura[i]
	}

	mediaAlt = somaAlt/float64(len(listaAltura))
	porcMenos40kg = (menos40Kg/totalPessoas)*100

	fmt.Printf("Quantidade de pessoas acima dos 50 anos: %.0f\nMédia das alturas das pessoas de 10 e 20 anos: %.2f\nPorcentagem de pessoas com peso inferior a 40kg: %.0f", mais50, mediaAlt, porcMenos40kg)

}