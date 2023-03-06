package main

import (
	"fmt"
	"os"
)

func main() {
	// Funcion
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic ocurred:", err)
		}
	}()

	if file, err := os.Open("hola.txt"); err != nil {
		panic("No fue posible obtener el archivo!!")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo")
			file.Close()
		}()
		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}

}
