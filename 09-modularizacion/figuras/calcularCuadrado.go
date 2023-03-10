package figuras

type Cuadrado struct {
	Lado float64
}

func (p *Cuadrado) area() float64 {
	return p.Lado * p.Lado
}

func (p *Cuadrado) perimetro() float64 {
	return 4 * p.Lado
}
