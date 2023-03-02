package main

import "fmt"

func modicarValor(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "Hola desde la funcion"
}

func main() {
	cadena := "Hola mundo de Go"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la funcion", cadena)

	modicarValor(&cadena) // Referencia

	fmt.Println("Despues de la funcion", cadena)
}
