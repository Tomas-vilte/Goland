package figuras

import "fmt"

type Geometrica interface {
	area() float64
	perimetro() float64
}

func Medidas(g Geometrica) {
	fmt.Println("Medidas: ", g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimetro", g.perimetro())
}
