package main

import (
	"fmt"
)

func somar(numero1 int, numero2 int) int {
	return numero1 + numero2
}

func CalculoMatematico(numero1, numero2 int) (int, int) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2
	return soma, subtracao
}

func main() {
	var numeroRetorno int = somar(10, 20)
	fmt.Println(numeroRetorno)

	var f = func(texto string) {
		fmt.Println(texto)
	}
	f("Texto da função f")

	resultadoSoma, resultadoSubtracao := CalculoMatematico(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma1, _ := CalculoMatematico(10, 15)
	fmt.Println(resultadoSoma1)

	_, resultadoSubtracao1 := CalculoMatematico(10, 15)
	fmt.Println(resultadoSubtracao1)
}
