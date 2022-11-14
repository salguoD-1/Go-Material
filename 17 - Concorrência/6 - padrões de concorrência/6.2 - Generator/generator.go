package main

import (
	"fmt"
	"time"
)

func main() {
	// A variável é um canal do tipo "recebe dados"
	canal := escrever("Olá, mundo!")

	// Chamamos o canal 10 vezes.
	for i := 0; i < 10; i++ {
		// Passamos o canal
		fmt.Println(<-canal)
	}
}

// Retorna um canal do tipo string
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}