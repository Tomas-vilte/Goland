package models

type Persona struct {
	nombre string
	edad   int
}

func (p *Persona) NewPerson(nombre string, edad int) {
	p.nombre = nombre
	p.edad = edad
}

func (p *Persona) GetName() string {
	return p.nombre
}

func (p *Persona) SetName(nombre string) {
	p.nombre = nombre
}

func (e *Persona) GetAge() int {
	return e.edad
}

func (e *Persona) SetAge(age int) {
	e.edad = age
}
