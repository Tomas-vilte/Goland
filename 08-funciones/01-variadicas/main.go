package main

import "fmt"

func sumar(nombre string, numeros ...int) (string, int) {
	msg := fmt.Sprintf("La suma de %s es: ", nombre)
	var total int
	for _, num := range numeros {
		total += num

	}
	return msg, total
}

func main() {
	mensaje, result := sumar("Thomas", 5, 5)
	fmt.Println(mensaje, result)
}
