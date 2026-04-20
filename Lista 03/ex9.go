package main

import "fmt"

var qntAlunos, qntAprovados, qntExame, qntReprovados int
var nota1, nota2, somaMedias, mediaClasse float64
var mediaNotas = []float64{}
var aprovacao string

func main() {
	fmt.Print("Informe a quantidade de alunos: ")
	fmt.Scan(&qntAlunos)

	for i:=0; i < qntAlunos; i++ {
		fmt.Printf("Informe as duas notas do aluno %d: ", i+1)
		fmt.Scan(&nota1, &nota2)
		mediaNotas = append(mediaNotas, (nota1+nota2)/2)
	}

	for i:=0; i < qntAlunos; i++ {
		if mediaNotas[i] <= 3 {
			aprovacao = "Reprovado"
			qntReprovados++
		} else
		if mediaNotas[i] > 3 && mediaNotas[i] < 7 {
			aprovacao = "Exame"
			qntExame++
		} else 
		if mediaNotas[i] >= 7 && mediaNotas[i] <= 10 {
			aprovacao = "Aprovado"
			qntAprovados++
		}

		somaMedias += mediaNotas[i]

		fmt.Printf("A média do aluno %d é %.2f. %s\n", i+1, mediaNotas[i], aprovacao)
	}

	mediaClasse = somaMedias/float64(len(mediaNotas))

	fmt.Printf("Total de alunos aprovados: %d\nTotal de alunos de exame: %d\nTotal de alunos reprovados: %d\nMédia da classe: %.2f", qntAprovados, qntExame, qntReprovados, mediaClasse)
}