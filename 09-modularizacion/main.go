package main

import (
	"fmt"
	"paquetes/models"
)

func main() {
	/*
		mensajes.Imprimir()
		mensajes.Hola()
		cua := figuras.Cuadrado{
			Lado: 8,
		}
		figuras.Medidas(&cua)

		cir := figuras.Circulo{
			Radio: 5,
		}
		figuras.Medidas(&cir)
	*/
	p1 := models.Persona{}
	p1.NewPerson("Lucas", 28)
	fmt.Println(p1.GetName())
	p1.SetName("Joan")

	fmt.Println(p1)
}
