package main

import f "fmt"

var hrs, min, seg, segTotais int

func main() {
	f.Print("Escreva as horas, minutos e segundos: ")
	f.Scan(&hrs, &min, &seg)

	segTotais = hrs*3600 + min*60 + seg

	f.Printf("O TEMPO EM SEGUNDOS E = %d", segTotais)
}