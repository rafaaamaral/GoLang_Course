package main

import (
	"fmt"
)

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float64
}

type Estudante struct {
	Pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := Pessoa{"João", "Pedro", 20, 1.75}
	fmt.Println(p1)
	p2 := Estudante{p1, "Engenharia", "USP"}

	fmt.Println(p2.nome)
	fmt.Println(p2.altura)
}
