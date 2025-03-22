package main

import "fmt"

func somar(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func escrever(texto string, numeros ...int) {

	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	soma := somar(1, 2, 3, 4, 5)
	fmt.Println(soma)

	escrever("Olá Mundo", 1, 2, 3, 4, 5)
}
