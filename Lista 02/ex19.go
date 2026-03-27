package main

import f "fmt"
import "math"

var figura int
var area, volume, raio, altura float64
var pi float64 = 3.14

func main() {
	f.Print("Qual o tipo da figura?(1-cone / 2-cilindro / 3-esfera): ")
	f.Scan(&figura)
	switch figura {
	case 1:
		f.Print("Qual o valor do raio e da altura do cone?")
		f.Scan(&raio, &altura)
		volume = (pi*raio*raio*altura)/3
		area = pi*raio*(math.Sqrt((raio*raio) + (altura*altura))) + pi*raio*raio
	case 2:
		f.Print("Qual o valor do raio e da altura do cilindro?")
		f.Scan(&raio, &altura)
		volume = pi*raio*raio*altura
		area = 2*pi*raio*raio + 2*pi*raio*altura
	case 3:
		f.Print("Qual o valor do raio da esfera?")
		f.Scan(&raio)
		volume = (4*pi*raio*raio*raio)/3
		area = 4*pi*raio*raio
	}

	f.Printf("Area = %.2f \nVolume = %.2f", area, volume)
}

