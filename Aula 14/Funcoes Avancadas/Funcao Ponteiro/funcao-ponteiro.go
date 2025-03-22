package main

import "fmt"

func inversorSinal(numero int) int {
	return numero * -1
}

func inversorSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	fmt.Println("Função Ponteiro")

	numero := 20
	numeroInvertido := inversorSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	inversorSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
