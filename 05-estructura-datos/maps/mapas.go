package main

import "fmt"

func main() {
	dias := make(map[int]string)
	fmt.Println(dias)

	// Agregar datos
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"
	dias[7] = "Domingo"
	fmt.Println(dias)

	dias[7] = "SABADO"
	delete(dias, 1)
	fmt.Println(dias, len(dias))

	// Nuevo mapa
	estudiantes := make(map[string][]int)
	estudiantes["Alex"] = []int{13, 15, 17}
	estudiantes["Roel"] = []int{14, 13, 17}

	fmt.Println(estudiantes)

	fmt.Println(estudiantes["Alex"][0])
}
