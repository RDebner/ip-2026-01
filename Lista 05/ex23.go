package main

import "fmt"

var janela = [24]int{}
var corredor = [24]int{}
var janOuCor string
var acentosDispJanela, acentosDispCorredor = []int{}, []int{}

func main() {
	fmt.Print("Preencha os acentos da janela: ")
	for i := 0; i < 24; i++ {
		fmt.Scan(&janela[i])
		if janela[i] == 0 {
				acentosDispJanela = append(acentosDispJanela, i+1)
		}
	}
	fmt.Print("Preencha os acentos do corredor: ")
	for i := 0; i < 24; i++ {
		fmt.Scan(&corredor[i])
		if corredor[i] == 0 {
				acentosDispCorredor = append(acentosDispCorredor, i+1)
		}
	}

	if len(acentosDispJanela) == 0 && len(acentosDispCorredor) == 0 {
		fmt.Print("Não há acentos disponíveis no ônibus.")
	} else {
		fmt.Print("Deseja poltrona na janela ou no corredor? ")
		fmt.Scan(&janOuCor)
	
		if janOuCor == "Janela" || janOuCor == "janela" {
			if len(acentosDispJanela) > 0 {
				fmt.Printf("Acentos disponíveis na janela: ")
				for i := 0; i < len(acentosDispJanela); i++ {
					if i == len(acentosDispJanela)-1 {
						fmt.Printf("%d",acentosDispJanela[i])
					} else {
						fmt.Printf("%d, ",acentosDispJanela[i])
					}
				}
			} else {
				fmt.Print("Não há acentos disponíveis na janela")
			}
		} else
		if janOuCor == "Corredor" || janOuCor == "corredor" {
			if len(acentosDispCorredor) > 0 {
				fmt.Printf("Acentos disponíveis na corredor: ")
				for i := 0; i < len(acentosDispCorredor); i++ {
					if i == len(acentosDispCorredor)-1 {
						fmt.Printf("%d",acentosDispCorredor[i])
					} else {
						fmt.Printf("%d, ",acentosDispCorredor[i])
					}
				}
			} else {
				fmt.Print("Não há acentos disponíveis no corredor")
			}
		}
	}


	// fmt.Print(janela)
	// fmt.Print(corredor)
}
