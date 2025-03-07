package main

import (
	. "fmt"
)

type Usuario struct {
	nome     string
	idade    int
	endereco Endereco
}

type Endereco struct {
	rua    string
	numero int
}

func main() {
	var usuario Usuario
	usuario.nome = "Nome do Usuário"
	usuario.idade = 30

	var endereco Endereco
	endereco.rua = "Rua do Usuário"
	endereco.numero = 10
	usuario.endereco = endereco
	Println(usuario)

	usuario2 := Usuario{"Nome do Usuário 2", 25, Endereco{"Rua do Usuário 2", 10}}
	Println(usuario2)

	usuario3 := Usuario{idade: 35}
	usuario3.nome = "Nome do Usuário 3"
	Println(usuario3)
}
