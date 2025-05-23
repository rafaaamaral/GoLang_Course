package main

import "fmt"

func main() {
	canal := make(chan string, 2) // Canal com buffer de 2 mensagens

	canal <- "Olá Mundo!"   // Envia mensagem para o canal
	canal <- "Olá Mundo 2!" // Envia outra mensagem para o canal

	mensagem1 := <-canal // Recebe a primeira mensagem do canal
	mensagem2 := <-canal // Recebe a segunda mensagem do canal

	fmt.Println(mensagem1) // Imprime a primeira mensagem
	fmt.Println(mensagem2) // Imprime a segunda mensagem
}
