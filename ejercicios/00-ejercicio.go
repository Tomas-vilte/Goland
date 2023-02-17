package main

import "fmt"

//Practica 01: Suma de dos números enteros
//Enunciado: dado dos números enteros, hallar la suma.

//Análisis: para la solución de este problema, se requiere que el usuario ingrese dos números enteros
//y el sistema realice el cálculo respectivo para hallar la suma.

func suma() {
	// Declaro la variable numb1, numb2 de tipo "int" para almacenar los dos numeros enteros
	// que ingrese el usuario
	var num1, num2 int
	fmt.Print("Ingrese el primer número: ")
	// Scan, para leer los valores ingresados por el usuario y almacenarlos en las variables
	fmt.Scan(&num1)
	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&num2)
	// Creo una varible llamada suma donde almacena la suma de numb1 y numb2
	suma := num1 + num2
	fmt.Println("La suma de", num1, "+", num2, "es:", suma)
}

func main() {
	suma()
}
