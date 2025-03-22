package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) salvar() {
	fmt.Println(p.Nome, "salvo")
}

func (p *Pessoa) fezAniversario() {
	p.Idade++
}

func main() {
	pessoa1 := Pessoa{"João", 20}
	pessoa2 := Pessoa{"Maria", 30}

	pessoa1.salvar()
	pessoa2.salvar()

	pessoa2.fezAniversario()
	fmt.Println(pessoa2.Idade)
}
