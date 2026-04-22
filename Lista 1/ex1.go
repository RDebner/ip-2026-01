package main

import "fmt"

var nota1, nota2, nota3, media float64

func main() {
	fmt.Print("Escreva as notas: ")
	fmt.Scan(&nota1, &nota2, &nota3)

	media = (nota1 + nota2 + nota3) / 3

	if media >= 6 {
		fmt.Printf("MEDIA = %.2f\nAPROVADO", media)
	} else 
	if media < 6 {
		fmt.Printf("MEDIA = %.2f\nREPROVADO", media)
	}

}
