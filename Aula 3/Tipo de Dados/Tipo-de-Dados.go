package main

import (
	"errors"
	"fmt"
)

func main() {
	//Tipo de Dados int pode ser 8, 16, 32 ou 64 bits
	var numero int = 1000000000000000000
	fmt.Println(numero)

	//Tipo de Dados uint pode ser 8, 16, 32 ou 64 bits
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//Tipo de Dados rune é um alias para int32
	var numero3 rune = 123456
	fmt.Println(numero3)

	//Tipo de Dados byte é um alias para uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	//Tipo de Dados float32 e float64
	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 123000.45
	fmt.Println(numeroReal2)

	//Tipo de Dados string
	var texto string = "Texto"
	fmt.Println(texto)

	//Tipo de Dados boolean
	var booleano bool = true

	if booleano {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")

	}

	//Tipo de Dados Erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
