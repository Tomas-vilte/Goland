package main

import "fmt"

func main() {
	a := 4
	b := 5
	var r bool

	// Igualdad
	r = a == b
	fmt.Printf("%d Es igual que el segundo valor %d? %t \n", a, b, r)

	// Distinto
	r = a != b
	fmt.Printf("%d Es distinto que %d? %t \n", a, b, r)

	// Mayor que
	r = a > b
	fmt.Printf("%d Es mayor que %d? %t \n", a, b, r)

	// Menor que
	r = a < b
	fmt.Printf("%d Es menor que %d? %t \n", a, b, r)

	// Mayor o igual que
	r = a >= b
	fmt.Printf("%d Es mayor o igual que %d? %t \n", a, b, r)

	// Menor o igual que
	fmt.Printf("%d Es menor o igual que %d? %t \n", a, b, r)
}
