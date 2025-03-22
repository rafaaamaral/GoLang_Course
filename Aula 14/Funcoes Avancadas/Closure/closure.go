package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"
	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro do main"
	fmt.Println(texto)

	funcaonova := closure()
	funcaonova()
}
