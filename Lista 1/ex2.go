package main

import "fmt"

var qntJogos int
var categPop, categGer, categArq, categCad, qntPessoas, renda float64

func main() {
	fmt.Print("Informe a quantidade de jogos: ")
	fmt.Scan(&qntJogos)

	var jogos []float64

	for i := 0; i < qntJogos; i++ {
		fmt.Println("Informe a quantidade de pessoas no jogo", i + 1, "e as porcentagens em cada categoria: ")
		fmt.Scan(&qntPessoas, &categPop, &categGer, &categArq, &categCad)
		calcularRenda()
		jogos = append(jogos, renda)
	}

	for i := 0; i < qntJogos ; i++ {
		fmt.Printf("A RENDA DO JOGO N.%d E: %.2f\n", i + 1, jogos[i] )
	}
}

func calcularRenda() {

	var qntPop, qntGer, qntArq, qntCad float64
	
	qntPop = qntPessoas * categPop/100
	qntGer = qntPessoas * categGer/100
	qntArq = qntPessoas * categArq/100
	qntCad = qntPessoas * categCad/100

	renda = (qntPop*1 + qntGer*5 + qntArq*10 + qntCad*20)
}


// 55000 20.2 50.4 30.2 10.2
// 49732 15.2 53.4 20.24 11.16
// 67890 30.0 42.20 23.8 4.0