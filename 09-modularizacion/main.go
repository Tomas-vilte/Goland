package main

import (
	"paquetes/figuras"
	"paquetes/mensajes"
)

func main() {
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

}
