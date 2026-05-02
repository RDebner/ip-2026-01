package main

import "fmt"

type Funcionario struct {
	numero   int
	numMeses int
}

var funcionario = Funcionario{}
var listaFuncionarios = []Funcionario{}

func main() {

	for {
		fmt.Print("Informe o número e a quantidade de meses do funcionário: ")
		fmt.Scan(&funcionario.numero, &funcionario.numMeses)

		if funcionario.numero == 0 && funcionario.numMeses == 0 {
			break
		}

		listaFuncionarios = append(listaFuncionarios, funcionario)
	}

	var funcionariosRecentes = [3]Funcionario{}
	funcionariosRecentes[0] = listaFuncionarios[0]
	funcionariosRecentes[1] = listaFuncionarios[1]
	funcionariosRecentes[2] = listaFuncionarios[2]

	for i := 3; i < len(listaFuncionarios); i++ {
		ordenar(&funcionariosRecentes)

		if listaFuncionarios[i].numMeses < funcionariosRecentes[2].numMeses {
			funcionariosRecentes[2] = listaFuncionarios[i]
		}
	}

	ordenar(&funcionariosRecentes)

	fmt.Printf("Os funcionários mais recentes são: %d, %d, %d\n", funcionariosRecentes[0].numero, funcionariosRecentes[1].numero, funcionariosRecentes[2].numero)
}

func ordenar(funci *[3]Funcionario) {
	if funci[0].numMeses > funci[1].numMeses {
		funci[0], funci[1] = funci[1], funci[0]
	}
	if funci[0].numMeses > funci[2].numMeses {
		funci[0], funci[1], funci[2] = funci[2], funci[0], funci[1]
	}
	if funci[1].numMeses > funci[2].numMeses {
		funci[1], funci[2] = funci[2], funci[1]
	}
}
