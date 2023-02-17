package main

import "fmt"

func main() {
	a := 20
	b := 10

	// Suma
	result := a + b
	fmt.Println("La suma entre a+b es:", result)

	// Resta
	result = a - b
	fmt.Println("La resta entre a-b es:", result)

	// Multiplicacion
	result = a * b
	fmt.Println("La multiplicacion entre a*b es:", result)

	// Division
	var div float64 = 3.0 / 2.0
	fmt.Println("La division entre a/2 es:", div)

	// Modulo
	result = a % b
	fmt.Println("El resto de la division es:", result)
}
