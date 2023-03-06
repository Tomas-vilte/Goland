package main

import (
	"errors"
	"fmt"
)

func division(dividiendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir entre 0!!")
	} else {
		return dividiendo / divisor, nil
	}
}

func main() {
	result, err := division(4, 1)
	if err == nil {
		fmt.Println("El resultado de la division es: ", result)
	} else {
		fmt.Println(err)
	}
}
