package main

import "fmt"

func main() {
	// arrays
	var numeros [5]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println(numeros)
	fmt.Println(numeros[4])

	// arrays con valores
	nombres := [3]string{"Thomas", "Joan", "Emilse"}
	fmt.Println(nombres)

	colores := [...]string{
		"Rojo",
		"Azul",
		"Amarillo",
	}
	fmt.Println(colores, len(colores))

	// Indices definidos
	monedas := [...]string{
		0: "Dolar",
		1: "Pesos",
		2: "Real",
		3: "Euros",
	}
	fmt.Println(monedas, len(monedas))

	// sub Array
	subMoneda := monedas[0:3]
	fmt.Println(subMoneda)
}
