package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo!")

	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}

	fmt.Println("Fim do Programa")
}

func escrever(mensagem string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor Recebido: %s", mensagem)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	return canal
}
