package main

import "fmt"

func main() {
	//Not
	fmt.Println(!true)
	//And
	fmt.Println(true && true)
	//Or
	fmt.Println(true || true)

	b := 2
	r := b == 2 || !(4 > b)
	fmt.Println(r)
}
