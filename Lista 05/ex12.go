package main

import "fmt"

type Nota struct {
	nota float64
	freqAbs float64
	freqRel float64
}

var n = Nota{}
var notas = []Nota{}

func main() {

	for i := 0; i < 15 ; i++ {
		fmt.Scan(&n.nota)
		if n.nota >= 0 && n.nota <= 10 {
			repeticao := 0
			n.freqAbs = 1
			n.freqRel = n.freqAbs / float64(len(notas))
			for j := 0; j < len(notas); j++ {
				if n.nota == notas[j].nota {
					repeticao += 1
					notas[j].freqAbs += 1
					notas[j].freqRel = notas[j].freqAbs / 15
				} 
			}
			if repeticao != 1 {
			notas = append(notas, n)
			}
		} else {
			notaInválida(i)
		}
	}

	fmt.Println("Nota | Frequência Absoluta | Frequência Relativa")
	for i:=0; i < len(notas); i++ {
		fmt.Printf("%.2f | %.2f | %.2f\n", notas[i].nota, notas[i].freqAbs, notas[i].freqRel)
	}
}

func notaInválida(i int) {
		fmt.Print("Nota inválida, informe uma nota válida (0 - 10): ")
		fmt.Scan(&n.nota)
		if n.nota >= 0 && n.nota <= 10 {
			notas = append(notas, n)
		} else {
			notaInválida(i)
		}
}