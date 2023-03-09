package main

import (
	"fmt"
	"math"
)

/*
En esta práctica vamos a sacar área y perímetro de un cuadrado y círculo utilizando interfaces.

Cuadrado

área = ancho * altura

perimetro = 2*ancho + 2*altura

Círculo

área = pi * (radio * radio)

perimetro = 2*pi * radio
*/

type Geometrica interface {
	area() float64
	perimetro() float64
}

type Cuadrado struct {
	lado float64
}

type Circulo struct {
	radio float64
}

func (p *Cuadrado) area() float64 {
	return p.lado * p.lado
}

func (p *Cuadrado) perimetro() float64 {
	return 4 * p.lado
}

func (cir *Circulo) area() float64 {
	return math.Pi * (cir.radio * cir.radio)
}

func (cir *Circulo) perimetro() float64 {
	return 2 * math.Pi
}

func medidas(g Geometrica) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	cuadrado := Cuadrado{lado: 4}
	circulo := Circulo{radio: 5}
	medidas(&cuadrado)
	medidas(&circulo)

}
