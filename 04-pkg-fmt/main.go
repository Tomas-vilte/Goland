package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"
	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Thomas"
	edad := 19
	fmt.Printf("Mi nombre es %s y tengo %d \n", nombre, edad)

	mensaje := fmt.Sprintf("Mi nombre es %s y tengo %d", nombre, edad)

	fmt.Println(mensaje)

	fmt.Printf("Nombre: %T edad: %T \n", nombre, edad)
	fmt.Print("Ingrese su nombre:")
	fmt.Scanln(&nombre)

	fmt.Println("Otro nombre:", nombre)
}
