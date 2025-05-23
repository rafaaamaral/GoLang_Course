package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	for {
		mensagem, aberto := <-canal // Recebe o valor do canal

		if !aberto { // Verifica se o canal está fechado
			break // Sai do loop se o canal estiver fechado
		}

		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal) // Fecha o canal após o envio de todas as mensagens
}
