package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)
}
