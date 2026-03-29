package main

import "fmt"

var qntTemp int
var temp, tempCelsius float64
var listTempFharenheit []float64
var listTempCelsius []float64

func main() {
	fmt.Print("Quantas temperaturas vão ser convertidas? ")
	fmt.Scan(&qntTemp)

	for i := 0 ; i < qntTemp ; i++ {
		fmt.Println("Escreva a temperatura",i + 1,"em fharenheit: ")
		fmt.Scan(&temp)
		listTempFharenheit = append(listTempFharenheit, temp)
		converterTemp()
		listTempCelsius = append(listTempCelsius, tempCelsius)
	}

	for i := 0 ; i < qntTemp; i++ {
		fmt.Printf("%.2f FHARENHEIT EQUIVALE A %.2f CELSIUS\n", listTempFharenheit[i], listTempCelsius[i])
	}
}

func converterTemp() {
	tempCelsius = 5*(temp - 32) / 9
}