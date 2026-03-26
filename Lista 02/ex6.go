package main

import f "fmt"

var num1, num2 int

func main() {
	f.Print("Escreva o dividendo e o divisor: ")
	f.Scan(&num1, &num2)
	if num1 % num2 == 0 {
		f.Printf("O número %d é divisível por %d\n", num1, num2)
	} else { 
		f.Printf("O número %d não é divisível por %d\n", num1, num2)
	}
}