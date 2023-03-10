package main

import "fmt"

func saludar(nombre string) {
	//fmt.Println("Hola mundo")
	fmt.Println("Hola", nombre)
}

func sumar(numb1, numb2 int) int {
	//suma := numb1 + numb2
	//result := suma
	//fmt.Println("La suma entre los valores numb1 y numb2 es:", result)
	return numb1 + numb2

}
func main() {
	saludar("Thomas")
	result := sumar(5, 5)
	fmt.Println("La suma entre numb1 y numb2 es:", result)
}
