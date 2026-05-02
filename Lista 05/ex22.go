package main

import "fmt"

var contas = [10]int{}
var saldos = [10]float64{}
var conta, digito, ind int
var deposito, saque, ativoBanc float64

func main() {
	for i := 0; i < len(contas); i++ {
		fmt.Print("Informe o código da conta: ")
		fmt.Scan(&conta)
		verificarConta(conta, i)
		fmt.Print("Informe o saldo da conta: ")
		fmt.Scan(&saldos[i])
	}

	for {

		fmt.Println("\n1. Efetuar depósito\n2. Efetuar saque\n3. Consultar o ativo bancário\n4. Consultar saldo\n5. Finalizar o programa")

		fmt.Scan(&digito)

		if digito == 1 {
			fmt.Print("Informe o código da conta: ")
			fmt.Scan(&conta)
			contaEncontrada := false
			for i := 0; i < len(contas); i++ {
				if conta == contas[i] {
					ind = i
					contaEncontrada = true
				}
			}
			if contaEncontrada {
				fmt.Print("Informe o valor do depósito: ")
				fmt.Scan(&deposito)
				saldos[ind] += deposito
			} else {
				fmt.Print("\nConta não encontrada\n")
			}
		} else
		if digito == 2 {
			fmt.Print("Informe o código da conta: ")
			fmt.Scan(&conta)
			contaEncontrada := false
			for i := 0; i < len(contas); i++ {
				if conta == contas[i] {
					ind = i
					contaEncontrada = true
				}
			}
			if contaEncontrada {
				fmt.Print("Informe o valor do saque: ")
				fmt.Scan(&saque)
				if saque > saldos[ind] {
					fmt.Print("\nSaldo insuficiente\n")
				} else {
					saldos[ind] -= saque
				}
			} else {
				fmt.Print("\nConta não encontrada\n")
			}
		} else
		if digito == 3 {
			ativoBanc = 0
			for i:=0; i < len(saldos); i++ {
				ativoBanc += saldos[i]
			}
			fmt.Printf("\nAtivo Bancário: %.2f\n", ativoBanc)
		} else
		if digito == 4 {
			fmt.Print("Informe o código da conta: ")
			fmt.Scan(&conta)
			contaEncontrada := false
			for i := 0; i < len(contas); i++ {
				if conta == contas[i] {
					ind = i
					contaEncontrada = true
				}
			}
			if contaEncontrada {
				fmt.Printf("\nO saldo da conta %d é: %.2f\n", conta, saldos[ind])
			} else {
				fmt.Print("\nConta não encontrada\n")
			}
		} else 
		if digito == 5 {
			break
		}
	}
}

func verificarConta(c, ind int) {
	adicionarConta := true
	for i := 0; i < len(contas); i++ {
		if c == contas[i] {
			adicionarConta = false
			fmt.Print("Já existe uma conta com esse código, informe outro: ")
			fmt.Scan(&conta)
			verificarConta(conta, ind)
		}
	}

	if adicionarConta {
		contas[ind] = conta
	}
}
