package main

import "fmt"

func factorial(numero int) int {
	if numero == 0 {
		return 1
	} else {
		return numero * factorial(numero-1)
	}
}

func main() {
	var numero = 3
	fmt.Println("Factorial: ", factorial(numero))
}
