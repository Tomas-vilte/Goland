package main

import "fmt"

//Practica 02: Calcular cociente y el Residuo de dos números enteros
//Enunciado: halar el cociente y el residuo (resto) de dos números enteros.

//Análisis: para la solución de este problema, se requiere que el usuario ingrese dos números enteros
//y el sistema realice el cálculo respectivo para hallar el cociente y residuo.

func calcular() {
	fmt.Print("Ingrese un numero: ")
	// Declaracion de variables
	var numb1, numb2, cociente, resto int
	// Lectura de valores ingresados por el usuario
	fmt.Scanln(&numb1)
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scanln(&numb2)
	// Cálculo del cociente y el residuo
	cociente = numb1 / numb2
	resto = numb1 % numb2
	// Impresión de resultados
	fmt.Println("El cociente de", numb1, "entre", numb2, "es:", cociente)
	fmt.Println("El residuo de", numb1, "entre", numb2, "es:", resto)
}

func main() {
	calcular()
}
