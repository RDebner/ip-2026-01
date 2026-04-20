package main

import (
	"fmt"
)

var cpf, somaCpf int
var primeiroDigito, segundoDigito bool

func main() {
	fmt.Print("Informe um CPF para ser validado: ")
	fmt.Scan(&cpf)
	
	var digitos = []int{}

	for cpf > 0 { 
		digito := cpf % 10
		digitos = append(digitos, digito)
		cpf = cpf / 10
	}
	
	j := 2
	for i:=2; i < len(digitos); i++ {
		somaCpf += digitos[i]*j
		j++
	}

	restoSomaCpf := somaCpf%11

	if restoSomaCpf < 2 && digitos[1] == 0 {
		primeiroDigito = true
	} else 
	if restoSomaCpf >= 2 && digitos[1] == 11-restoSomaCpf {
		primeiroDigito = true
	} else {
		primeiroDigito = false
	}

	somaCpf = 0
	
	j = 2
	for i:=1; i < len(digitos); i++ {
		somaCpf += digitos[i]*j
		j++
	}
	
	restoSomaCpf = somaCpf%11

	if restoSomaCpf < 2 && digitos[0] == 0 {
		segundoDigito = true
	} else 
	if restoSomaCpf >= 2 && digitos[0] == 11-restoSomaCpf {
		segundoDigito = true
	} else {
		segundoDigito = false
	}
	
	if primeiroDigito && segundoDigito {
		fmt.Print("CPF válido")
	} else {
		fmt.Print("CPF inválido")
	}
}