package main

import "fmt"

// Variable global
var mensaje string

func funcion1() {
	mensaje = "Hola desde la funcion uno!"
	fmt.Println(mensaje)
}

func funcion2() {
	mensaje = "Hola desde la funcion dos!"
	fmt.Println(mensaje)
}
func main() {
	mensaje = "Hola desde la funcion main!"
	fmt.Println(mensaje)

	defer funcion1()
	funcion2()
	fmt.Println(mensaje)
}
