package main

import (
	"fmt"
)

func main() {
	if nombre, edad := "Thomas1", 19; nombre == "Thomas" {
		fmt.Println("Hola,", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad:%d\n", nombre, edad)
	}
	users := make(map[string]string)
	users["Alex"] = "Alex@email.com"
	users["Joan"] = "Joan@email.com"
	users["Alexis"] = "Alexis@email.com"

	// correo, ok := users["Alex"]
	if correo, ok := users["Alex"]; ok {
		fmt.Println(correo, ok)
	} else {
		fmt.Println("No fue posible encontrar la clave")
	}

}
