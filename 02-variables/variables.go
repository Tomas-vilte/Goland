package main

import "fmt"

func main() {
	var nombre1 string
	nombre1 = "Thomas"

	var apellido string = "Vilte"

	edad := 19

	fmt.Printf("Mi nombre es %s %s y tengo %d años", nombre1, apellido, edad)

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	const pi = 3.141592
	fmt.Println("El valor de pi es:", pi)
}
