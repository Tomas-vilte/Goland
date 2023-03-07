package main

import "fmt"

// Struct persona

type Persona struct {
	nombre string
	edad   int
}

// Metodo

func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad:%d\n", p.nombre, p.edad)
}

func (p *Persona) editarNombre(nombre string) {
	p.nombre = nombre
}

// Herencia
type Empleado struct {
	Persona
	sueldo float64
}

func main() {
	persona1 := Persona{"Thomas", 19}
	persona1.imprimir()
	//fmt.Println(persona1)
	persona1.editarNombre("Alexis")
	persona1.imprimir()
	persona2 := Persona{
		nombre: "Joan",
		edad:   25,
	}
	//fmt.Println(persona2)
	persona2.imprimir()
	persona2.editarNombre("Emilse")
	persona2.imprimir()

	emp1 := Empleado{sueldo: 60.656}
	emp1.nombre = "Lucas"
	emp1.edad = 29
	emp1.imprimir()
	fmt.Println(emp1)
}
