package mensajes

import "fmt"

func Hola() {
	fmt.Println("Hola desde el paquete mensaje")
}

func funcionPrivada() {
	fmt.Println("Hola desde mi funcion privada")
}

const mensaje = "Hola desde mi constante"

func Imprimir() {
	fmt.Println(mensaje)
	funcionPrivada()
}
