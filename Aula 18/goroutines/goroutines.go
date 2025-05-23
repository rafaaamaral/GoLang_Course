package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!")
	escrever("Olá Mundo 2!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
