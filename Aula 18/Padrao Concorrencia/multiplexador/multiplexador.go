package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo!"), escrever("Olá Go!"))
	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}

}

func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalDeSaida <- mensagem
			case mensagem := <-canal2:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}

func escrever(mensagem string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor Recebido: %s", mensagem)
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		}
	}()

	return canal
}
