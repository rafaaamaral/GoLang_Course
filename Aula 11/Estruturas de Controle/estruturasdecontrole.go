package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if condition := numero > 15; condition {
		fmt.Println("Maior que 15")
	}
}
